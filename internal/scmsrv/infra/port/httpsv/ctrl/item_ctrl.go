package ctrl

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/thanhpp/scm/internal/scmsrv/app"
	"github.com/thanhpp/scm/internal/scmsrv/infra/port/httpsv/dto"
	"github.com/thanhpp/scm/pkg/ginutil"
	"github.com/thanhpp/scm/pkg/logger"
)

type ItemCtrl struct {
	itemHandler app.ItemHandler
}

func NewItemCtrl(itemHandler app.ItemHandler) *ItemCtrl {
	return &ItemCtrl{
		itemHandler: itemHandler,
	}
}

func (ctrl ItemCtrl) CreateItem(c *gin.Context) {
	req := new(dto.CreateItemReq)

	if err := c.ShouldBind(req); err != nil {
		ginutil.RespErr(c, http.StatusNotAcceptable, err, ginutil.WithData(req))
		return
	}

	form, err := c.MultipartForm()
	if err != nil {
		ginutil.RespErr(c, http.StatusNotAcceptable, errors.WithMessage(err, "form multipart"))
		return
	}

	images := form.File["images"]

	newItem, err := ctrl.itemHandler.CreateItem(
		c.Request.Context(), req.SKU, req.Name, req.Desc,
		req.ItemTypeID, images,
	)
	if err != nil {
		ginutil.RespErr(c, http.StatusInternalServerError, err)
		return
	}

	resp := new(dto.ItemInfoResp)
	resp.Set200OK()
	resp.SetData(newItem)

	c.JSON(http.StatusOK, resp)
}

func (ctrl ItemCtrl) GetList(c *gin.Context) {
	pagination := ginutil.NewPaginationQuery(c)

	items, err := ctrl.itemHandler.GetList(c.Request.Context(), pagination.Page, pagination.Size)
	if err != nil {
		ginutil.RespErr(c, http.StatusInternalServerError, err)
		return
	}
	logger.Debugw("Get list item", "items", items)

	resp := new(dto.GetListItemInfoResp)
	resp.Set200OK()
	resp.SetData(items)
	logger.Debugw("Get list item", "resp", resp.Data)

	c.JSON(http.StatusOK, resp)
}

func (ctrl ItemCtrl) CreateItemType(c *gin.Context) {
	req := new(dto.CreateItemTypeReq)

	if err := c.ShouldBind(req); err != nil {
		ginutil.RespErr(c, http.StatusNotAcceptable, err, ginutil.WithData(req))
		return
	}

	newItemType, err := ctrl.itemHandler.CreateItemType(c.Request.Context(), req.Name, req.Desc)
	if err != nil {
		ginutil.RespErr(c, http.StatusInternalServerError, err)
		return
	}

	resp := new(dto.ItemTypeInfoResp)
	resp.Set200OK()
	resp.SetData(newItemType)

	c.JSON(http.StatusOK, resp)
}

func (ctrl ItemCtrl) GetAllItemType(c *gin.Context) {
	itemTypes, err := ctrl.itemHandler.GetAllItemType(c.Request.Context())
	if err != nil {
		ginutil.RespErr(c, http.StatusInternalServerError, err)
		return
	}

	resp := new(dto.GetAllItemTypeResp)
	resp.Set200OK()
	resp.SetData(itemTypes)

	c.JSON(http.StatusOK, resp)
}
