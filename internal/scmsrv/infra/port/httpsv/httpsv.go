package httpsv

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/thanhpp/scm/internal/scmsrv/app"
	"github.com/thanhpp/scm/internal/scmsrv/infra/port/httpsv/auth"
	"github.com/thanhpp/scm/internal/scmsrv/infra/port/httpsv/ctrl"
	"github.com/thanhpp/scm/pkg/booting"
	"github.com/thanhpp/scm/pkg/configx"
	"github.com/thanhpp/scm/pkg/ginutil"
	"github.com/thanhpp/scm/pkg/logger"
)

type HTTPServer struct {
	cfg    configx.HTTPServerConfig
	app    *app.App
	jwtSrv auth.JWTSrv
}

func NewHTTPServer(
	cfg configx.HTTPServerConfig, app *app.App,
) *HTTPServer {
	httpServer := &HTTPServer{
		cfg:    cfg,
		app:    app,
		jwtSrv: auth.NewJWTSrvImpl(),
	}

	return httpServer
}

func (s *HTTPServer) Daemon() booting.Daemon {
	server := http.Server{
		Addr:    fmt.Sprintf("%s:%s", s.cfg.Host, s.cfg.Port),
		Handler: s.newRouter(),
	}

	return func(ctx context.Context) (start func() error, cleanup func()) {
		start = func() error {
			logger.Infof("starting server at %s...", server.Addr)
			if err := server.ListenAndServe(); err != nil {
				if errors.Is(err, http.ErrServerClosed) {
					return nil
				}
				return err
			}

			return nil
		}

		cleanup = func() {
			cancelCtx, cancel := context.WithTimeout(ctx, time.Second*5)
			defer cancel()

			if err := server.Shutdown(cancelCtx); err != nil {
				logger.Errorf("shutdown http server err %v", err)
			}
		}

		return start, cleanup
	}
}

func (s *HTTPServer) newRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"*"},
		AllowHeaders:     []string{"*"},
		ExposeHeaders:    []string{"*"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return true
		},
		MaxAge: 24 * time.Hour,
	}))

	// init
	importTicketCtrl := ctrl.NewImportTicket(s.app.ImportTicketHandler)
	supplierCtrl := ctrl.NewSupplier(s.app.SupplierHandler)
	storageCtrl := ctrl.NewStorageCtrl(s.app.StorageHandler)
	itemCtrl := ctrl.NewItemCtrl(s.app.ItemHandler)
	filesCtrl := ctrl.NewFileCtrl()
	userCtrl := ctrl.NewUserCtrl(s.app.UserHandler, auth.NewJWTSrvImpl())

	importTicketGr := r.Group("import_ticket")
	{
		importTicketGr.Use(s.authMiddleware())
		importTicketGr.GET("", importTicketCtrl.GetListImportTickets)
		importTicketGr.GET("/:id", importTicketCtrl.GetImportTicket)
		importTicketGr.GET("/:id/serial", importTicketCtrl.GetSerialsByImportTicketID)

		importTicketGr.POST("", importTicketCtrl.Create)
		importTicketGr.POST("serials", importTicketCtrl.GenSerial)
	}

	r.GET("serial/:seri", importTicketCtrl.GetSeriData)

	supplierGr := r.Group("supplier")
	{
		supplierGr.Use(s.authMiddleware())
		supplierGr.GET("", supplierCtrl.GetList)
		supplierGr.GET("/:id", supplierCtrl.GetSupplier)

		supplierGr.POST("", supplierCtrl.Create)
		supplierGr.PUT("/:id", supplierCtrl.Update)
	}

	storageGr := r.Group("storage")
	{
		storageGr.Use(s.authMiddleware())
		storageGr.GET("", storageCtrl.GetList)
		storageGr.GET("/:id", storageCtrl.Get)

		storageGr.POST("", storageCtrl.Create)
		storageGr.PUT("/:id", storageCtrl.Update)
	}

	itemGr := r.Group("item")
	{
		itemGr.Use(s.authMiddleware())
		itemGr.GET("", itemCtrl.GetList)
		itemGr.POST("", itemCtrl.CreateItem)
		itemGr.PUT("/:sku", itemCtrl.UpdateItem)
	}

	itemTypeGr := r.Group("item-type")
	{
		itemGr.Use(s.authMiddleware())
		itemTypeGr.GET("", itemCtrl.GetAllItemType)
		itemTypeGr.POST("", itemCtrl.CreateItemType)
		itemTypeGr.PUT("/:id", itemCtrl.UpdateItemType)
	}

	userGr := r.Group("user")
	{
		userGr.Use(s.authMiddleware())
		userGr.GET("", userCtrl.GetUsers)
		userGr.PATCH("/:id/password", userCtrl.UpdateUserPassword)
	}

	r.POST("signup", userCtrl.NewUser)
	r.POST("login", userCtrl.Login)

	r.GET("files/:filename", filesCtrl.ServeFile)

	return r
}

const (
	bearerPrefix       = "Bearer "
	bearerPrefixLength = len(bearerPrefix)
)

func (s *HTTPServer) authMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if len(authHeader) < bearerPrefixLength {
			ginutil.RespErr(c, http.StatusUnauthorized, errors.New("invalid length"))
			c.Abort()
			return
		}

		authHeader = authHeader[bearerPrefixLength:]

		_, err := s.jwtSrv.Validate(authHeader)
		if err != nil {
			ginutil.RespErr(c, http.StatusUnauthorized, err)
			c.Abort()
			return
		}

		c.Next()
	}
}
