package ctrl

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
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

func (ctrl StorageCtrl) Update(c *gin.Context) {
	id, err := getIDFromQuery(c)
	if err != nil {
		ginutil.RespErr(c, http.StatusNotAcceptable, err)
		return
	}

	req := new(dto.CreateStorageReq)
	if err := c.ShouldBind(req); err != nil {
		ginutil.RespErr(c, http.StatusNotAcceptable, err, ginutil.WithData(req))
		return
	}

	if err := ctrl.storageHandler.UpdateStorage(c.Request.Context(), id, req.Name, req.Desc, req.Location); err != nil {
		ginutil.RespErr(c, http.StatusInternalServerError, err, ginutil.WithData(err))
		return
	}

	ginutil.RespOK(c, nil)
}

func getIDFromQuery(c *gin.Context) (int, error) {
	strID := c.Query("id")

	iID, err := strconv.Atoi(strID)
	if err != nil {
		return 0, errors.WithMessage(err, "convert id to int")
	}

	return iID, nil
}
