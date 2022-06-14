package ctrl

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/thanhpp/scm/internal/scmsrv/app"
	"github.com/thanhpp/scm/internal/scmsrv/infra/port/httpsv/dto"
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
		c.AbortWithError(http.StatusNotAcceptable, err)
		return
	}

	newSupplier, err := ctrl.supplierHandler.Create(c.Request.Context(), req.Name, req.Email, req.Phone)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, newSupplier)
}
