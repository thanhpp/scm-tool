package ctrl

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/thanhpp/scm/internal/scmsrv/app"
	"github.com/thanhpp/scm/internal/scmsrv/infra/port/httpsv/dto"
	"github.com/thanhpp/scm/pkg/ginutil"
)

type ItemCtrl struct {
	itemHandler app.ItemHandler
}

func NewItemCtrl(itemHandler app.ItemHandler) *ItemCtrl {
	return &ItemCtrl{
		itemHandler: itemHandler,
	}
}

func (ctrl ItemCtrl) Create(c *gin.Context) {
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

	newItem, err := ctrl.itemHandler.Create(
		c.Request.Context(), req.SKU, req.Name, req.Desc,
		req.ItemTypeID, images,
	)
	if err != nil {
		ginutil.RespErr(c, http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, newItem)
}
