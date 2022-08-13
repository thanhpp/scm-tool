package ctrl

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/thanhpp/scm/internal/scmsrv/app"
	"github.com/thanhpp/scm/internal/scmsrv/infra/port/httpsv/auth"
	"github.com/thanhpp/scm/internal/scmsrv/infra/port/httpsv/dto"
	"github.com/thanhpp/scm/pkg/ginutil"
)

type UserCtrl struct {
	userHandler app.UserHandler
	jwt         auth.JWTSrv
}

func NewUserCtrl(handler app.UserHandler, jwt auth.JWTSrv) *UserCtrl {
	return &UserCtrl{
		userHandler: handler,
		jwt:         jwt,
	}
}

func (ctrl UserCtrl) NewUser(c *gin.Context) {
	req := new(dto.CreateUserReq)

	if err := c.ShouldBind(req); err != nil {
		ginutil.RespErr(c, http.StatusNotAcceptable, err, ginutil.WithData(req))
		return
	}

	newUser, err := ctrl.userHandler.CreateUser(c.Request.Context(), req.Name, req.Username, req.Password)
	if err != nil {
		ginutil.RespErr(c, http.StatusInternalServerError, err, ginutil.WithData(req))
		return
	}

	resp := new(dto.UserInfoResp)
	resp.Set200OK()
	resp.SetData(newUser)

	c.JSON(http.StatusOK, resp)
}

func (ctrl UserCtrl) Login(c *gin.Context) {
	req := new(dto.LoginReq)

	if err := c.ShouldBind(req); err != nil {
		ginutil.RespErr(c, http.StatusNotAcceptable, err, ginutil.WithData(req))
		return
	}

	user, err := ctrl.userHandler.ValidateUser(c, req.Username, req.Password)
	if err != nil {
		ginutil.RespErr(c, http.StatusInternalServerError, err)
		return
	}

	token, err := ctrl.jwt.GenToken(user)
	if err != nil {
		ginutil.RespErr(c, http.StatusInternalServerError, err)
		return
	}

	resp := new(dto.LoginResp)
	resp.Set200OK()
	resp.SetData(token.JWT, user)

	c.JSON(http.StatusOK, resp)
}

func (ctrl UserCtrl) GetUsers(c *gin.Context) {
	p := ginutil.NewPaginationQuery(c)

	users, err := ctrl.userHandler.GetUsers(c, p.Limit(), p.Offset())
	if err != nil {
		ginutil.RespErr(c, http.StatusNotAcceptable, err)
		return
	}

	resp := new(dto.RespGetUsers)
	resp.Set200OK()
	resp.SetData(users)

	c.JSON(http.StatusOK, resp)
}

func (ctrl UserCtrl) UpdateUserPassword(c *gin.Context) {
	id, err := getIDFromParam(c)
	if err != nil {
		ginutil.RespErr(c, http.StatusNotAcceptable, err)
		return
	}

	req := new(dto.ReqUpdateUserPass)
	if err := c.ShouldBindJSON(req); err != nil {
		ginutil.RespErr(c, http.StatusNotAcceptable, err)
		return
	}

	if err := ctrl.userHandler.UpdateUserPassword(c, id, req.NewPass); err != nil {
		ginutil.RespErr(c, http.StatusInternalServerError, err)
		return
	}

	ginutil.RespOK(c, nil)
}
