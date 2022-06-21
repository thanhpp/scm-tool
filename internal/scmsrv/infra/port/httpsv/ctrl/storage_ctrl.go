package ctrl

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/thanhpp/scm/internal/scmsrv/app"
	"github.com/thanhpp/scm/internal/scmsrv/infra/port/httpsv/dto"
	"github.com/thanhpp/scm/pkg/ginutil"
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
		ginutil.RespErr(c, http.StatusNotAcceptable, err, ginutil.WithData(req))
		return
	}

	newStorage, err := ctrl.storageHandler.Create(c.Request.Context(), req.Name, req.Desc, req.Location)
	if err != nil {
		ginutil.RespErr(c, http.StatusInternalServerError, err, ginutil.WithData(err))
		return
	}

	resp := new(dto.StorageInfoResp)
	resp.Set200OK()
	resp.SetData(newStorage)

	c.JSON(http.StatusOK, resp)
}
func (ctrl StorageCtrl) GetList(c *gin.Context) {
	pagination := ginutil.NewPaginationQuery(c)

	storages, err := ctrl.storageHandler.GetListStorages(c, pagination.Page, pagination.Size)
	if err != nil {
		ginutil.RespErr(c, http.StatusInternalServerError, err)
		return
	}

	resp := new(dto.GetListStoragesResp)
	resp.Set200OK()
	resp.SetData(storages)

	c.JSON(http.StatusOK, resp)
}
