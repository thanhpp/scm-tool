package ctrl

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/thanhpp/scm/internal/scmsrv/app"
	"github.com/thanhpp/scm/internal/scmsrv/infra/port/httpsv/dto"
	"github.com/thanhpp/scm/pkg/ginutil"
)

type SupplierCtrl struct {
	supplierHandler app.SupplierHanlder
}

func NewSupplier(handler app.SupplierHanlder) *SupplierCtrl {
	return &SupplierCtrl{
		supplierHandler: handler,
	}
}

func (ctrl SupplierCtrl) Create(c *gin.Context) {
	req := new(dto.CreateSupplierReq)

	if err := c.ShouldBind(req); err != nil {
		ginutil.RespErr(c, http.StatusNotAcceptable, err, ginutil.WithData(req))
		return
	}

	newSupplier, err := ctrl.supplierHandler.Create(c.Request.Context(), req.Name, req.Email, req.Phone)
	if err != nil {
		ginutil.RespErr(c, http.StatusInternalServerError, err, ginutil.WithData(req))
		return
	}

	resp := new(dto.SupplierInfoResp)
	resp.Set200OK()
	resp.SetData(newSupplier)

	c.JSON(http.StatusOK, resp)
}
