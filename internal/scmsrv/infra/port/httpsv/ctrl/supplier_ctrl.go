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

func (ctrl SupplierCtrl) GetList(c *gin.Context) {
	pagination := ginutil.NewPaginationQuery(c)

	suppliers, err := ctrl.supplierHandler.GetList(c, pagination.Page, pagination.Size)
	if err != nil {
		ginutil.RespErr(c, http.StatusInternalServerError, err)
		return
	}

	resp := new(dto.GetListSupplierResp)
	resp.Set200OK()
	resp.SetData(suppliers)

	c.JSON(http.StatusOK, resp)
}

func (ctrl SupplierCtrl) Update(c *gin.Context) {
	id, err := getIDFromParam(c)
	if err != nil {
		ginutil.RespErr(c, http.StatusNotAcceptable, err)
		return
	}

	req := new(dto.CreateSupplierReq)
	if err := c.ShouldBind(req); err != nil {
		ginutil.RespErr(c, http.StatusNotAcceptable, err, ginutil.WithData(req))
		return
	}

	if err := ctrl.supplierHandler.Update(c, id, req.Name, req.Email, req.Phone); err != nil {
		ginutil.RespErr(c, http.StatusNotAcceptable, err, ginutil.WithData(req))
		return
	}

	ginutil.RespOK(c, nil)
}
