package httpsv

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/thanhpp/scm/internal/nftsrv/app"
	"github.com/thanhpp/scm/internal/nftsrv/infra/port/httpsv/ctrl"
	"github.com/thanhpp/scm/pkg/booting"
	"github.com/thanhpp/scm/pkg/configx"
	"github.com/thanhpp/scm/pkg/logger"
)

type HTTPServer struct {
	cfg configx.HTTPServerConfig
	app *app.App
}

func NewHTTPServer(cfg configx.HTTPServerConfig, app *app.App) *HTTPServer {
	return &HTTPServer{
		cfg: cfg,
		app: app,
	}
}

func (s HTTPServer) Daemon() booting.Daemon {
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

func (s HTTPServer) newRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	ctrl := ctrl.NewNFTMinterCtrl(s.app)

	r.POST("mint", ctrl.MintNFT)

	r.GET("txhash/:txHash", ctrl.GetByTxHash)
	r.GET("tokenid/:tokenID", ctrl.GetByTokenID)

	return r
}
