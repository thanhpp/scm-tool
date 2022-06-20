package ctrl

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/thanhpp/scm/internal/scmsrv/app"
	"github.com/thanhpp/scm/internal/scmsrv/domain/entity"
	"github.com/thanhpp/scm/internal/scmsrv/infra/port/httpsv/dto"
	"github.com/thanhpp/scm/pkg/ginutil"
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
	req, err := dto.CustomBindCreateImportTicketReq(c)
	if err != nil {
		ginutil.RespErr(c, http.StatusNotAcceptable, err, ginutil.WithData(dto.CreateImportTicketReq{}))
		return
	}

	form, err := c.MultipartForm()
	if err != nil {
		ginutil.RespErr(c, http.StatusNotAcceptable, err, ginutil.WithData(req))
		return
	}

	billImages := form.File["bill_images"]
	productImages := form.File["product_images"]

	details := make([]entity.ImportTicketDetails, 0, len(req.Details))
	for i := range req.Details {
		detail, err := ctrl.importTickerHanlder.CreateImportDetails(
			c.Request.Context(), req.Details[i].ItemSKU, req.Details[i].BuyQuantity,
			0, req.Details[i].BuyPrice,
		)
		if err != nil {
			ginutil.RespErr(c, http.StatusNotAcceptable, err, ginutil.WithData(req.Details[i]))
			return
		}

		details = append(details, *detail)
	}

	importTicket, err := ctrl.importTickerHanlder.Create(
		c.Request.Context(), req.FromSupplierID, req.ToStorageID,
		req.SendTime, time.Time{}, req.Fee, details, billImages, productImages)
	if err != nil {
		ginutil.RespErr(c, http.StatusNotAcceptable, err)
		return
	}

	c.JSON(http.StatusOK, importTicket)
}

func (ctrl ImportTicketCtrl) GenSerial(c *gin.Context) {
	req := new(dto.GenSerialReq)

	if err := c.ShouldBind(req); err != nil {
		ginutil.RespErr(c, http.StatusNotAcceptable, err, ginutil.WithData(req))
		return
	}

	serials, err := ctrl.importTickerHanlder.GenSerials(c.Request.Context(), req.ImportTicketID)
	if err != nil {
		ginutil.RespErr(c, http.StatusInternalServerError, err, ginutil.WithData(req))
		return
	}

	resp := new(dto.GenSerialResp)
	resp.Set200OK()
	resp.SetData(serials)

	c.JSON(http.StatusOK, resp)
}
