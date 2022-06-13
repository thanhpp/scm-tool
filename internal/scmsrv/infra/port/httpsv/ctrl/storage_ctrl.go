package ctrl

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/thanhpp/scm/internal/scmsrv/app"
	"github.com/thanhpp/scm/internal/scmsrv/infra/port/httpsv/dto"
)

type StorageCtrl struct {
	storageHandler app.StorageHandler
}

func NewStorageCtrl(handler app.StorageHandler) *StorageCtrl {
	return &StorageCtrl{
		storageHandler: handler,
	}
}

func (ctrl StorageCtrl) Create(c *gin.Context) {
	req := new(dto.CreateStorageReq)

	if err := c.ShouldBind(req); err != nil {
		c.AbortWithError(http.StatusNotAcceptable, err)
		return
	}

	newStorage, err := ctrl.storageHandler.Create(c, req.Name, req.Desc, req.Location)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, newStorage)
}
