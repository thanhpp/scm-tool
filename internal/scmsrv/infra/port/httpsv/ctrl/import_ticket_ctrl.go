package ctrl

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/thanhpp/scm/internal/scmsrv/app"
	"github.com/thanhpp/scm/internal/scmsrv/infra/port/httpsv/dto"
)

type ImportTicketCtrl struct {
	importTickerHanlder app.ImportTicketHandler
}

func NewImportTicket(importTicketHandler app.ImportTicketHandler) *ImportTicketCtrl {
	return &ImportTicketCtrl{
		importTickerHanlder: importTicketHandler,
	}
}

func (ctrl ImportTicketCtrl) Create(c *gin.Context) {
	req := new(dto.CreateImportTicketReq)

	if err := c.ShouldBind(req); err != nil {
		c.AbortWithStatusJSON(http.StatusNotAcceptable, req)
		return
	}

	form, err := c.MultipartForm()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotAcceptable, req)
		return
	}

	billImages := form.File["bill_images"]
	productImages := form.File["product_images"]

	importTicket, err := ctrl.importTickerHanlder.Create(
		c, req.FromSupplierID, req.ToStorageID,
		req.SendTime, time.Time{}, req.Fee, nil, billImages, productImages)
	if err != nil {
		log.Println(err)
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, importTicket)
}
