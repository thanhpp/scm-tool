package dto

import (
	"encoding/json"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/thanhpp/scm/internal/scmsrv/domain/entity"
	"github.com/thanhpp/scm/pkg/ginutil"
)

type CreateImportTicketReq struct {
	FromSupplierID int                            `json:"from_supplier_id" form:"from_supplier_id"`
	ToStorageID    int                            `json:"to_storage_id" form:"to_storage_id"`
	SendTime       time.Time                      `json:"send_time" form:"send_time"`
	ReceiveTime    time.Time                      `json:"receive_time" form:"receive_time"`
	Fee            float64                        `json:"fee" form:"fee"`
	Details        []CreateImportTicketReqDetails `json:"-" form:"-"`
}

func CustomBindCreateImportTicketReq(c *gin.Context) (*CreateImportTicketReq, error) {
	req := new(CreateImportTicketReq)

	if err := c.ShouldBind(req); err != nil {
		return nil, errors.WithMessage(err, "parse create import ticket: bind")
	}

	if err := c.Request.ParseForm(); err != nil {
		return nil, errors.WithMessage(err, "parse create import ticket: parse form")
	}

	details := c.Request.FormValue("details")

	data, err := strconv.Unquote(details)
	if err != nil {
		return nil, errors.WithMessage(err, "parse create import ticket: unquote details")
	}

	if err := json.Unmarshal([]byte(data), &req.Details); err != nil {
		return nil, errors.WithMessage(err, "parse create import ticket: unmarshal details")
	}

	return req, nil
}

type CreateImportTicketReqDetails struct {
	ItemSKU     string  `json:"item_sku" form:"item_sku"`
	BuyQuantity int     `json:"buy_quantity" form:"buy_quantity"`
	BuyPrice    float64 `json:"buy_price" form:"buy_price"`
}

type GenSerialReq struct {
	ImportTicketID int `json:"import_ticket_id"`
}

type ImportTicketInfoRespData struct {
	ID                int       `json:"id"`
	FromSupplierID    int       `json:"from_supplier_id"`
	ToStorageID       int       `json:"to_storage_id"`
	Status            string    `json:"status"`
	SendTime          time.Time `json:"send_time"`
	ReceiveTime       time.Time `json:"receive_time"`
	Fee               float64   `json:"fee"`
	BillImagePaths    []string  `json:"bill_image_paths"`
	ProductImagePaths []string  `json:"product_image_paths"`
	Details           []struct {
		ItemSKU     string  `json:"item_sku"`
		BuyQuantity int     `json:"buy_quantity"`
		BuyPrice    float64 `json:"buy_price"`
	} `json:"details"`
}

func (d *ImportTicketInfoRespData) set(in *entity.ImportTicket) {
	d.ID = in.ID
	d.FromSupplierID = in.FromSupplier.ID
	d.ToStorageID = in.ToStorage.ID
	d.Status = in.Status.String()
	d.SendTime = in.SendTime
	d.ReceiveTime = in.ReceiveTime
	d.Fee = in.Fee

	for i := range in.BillImagePaths {
		d.BillImagePaths = append(d.BillImagePaths, buildFileURL(in.BillImagePaths[i]))
	}

	for i := range in.ProductImagePaths {
		d.ProductImagePaths = append(d.ProductImagePaths, buildFileURL(in.ProductImagePaths[i]))
	}

	for i := range in.Details {
		d.Details = append(d.Details, struct {
			ItemSKU     string  `json:"item_sku"`
			BuyQuantity int     `json:"buy_quantity"`
			BuyPrice    float64 `json:"buy_price"`
		}{
			ItemSKU:     in.Details[i].Item.SKU,
			BuyQuantity: in.Details[i].BuyQuantity,
			BuyPrice:    in.Details[i].BuyPrice,
		})
	}
}

type ImportTicketInfoResp struct {
	ginutil.RespTemplateError
	Data ImportTicketInfoRespData `json:"data"`
}

func (resp *ImportTicketInfoResp) SetData(in *entity.ImportTicket) {
	resp.Data.set(in)
}

type GenSerialRespData struct {
	ImportTicketID int      `json:"import_ticket_id"`
	ItemSKU        string   `json:"item_sku"`
	Serials        []string `json:"serials"`
}

type GenSerialResp struct {
	ginutil.RespTemplateError
	Data []*GenSerialRespData `json:"data"`
}

func (resp *GenSerialResp) SetData(serials []*entity.Serial) {
	m := make(map[string]*GenSerialRespData) // itemSKU - GenSerialRespData

	for i := range serials {
		if data, ok := m[serials[i].Item.SKU]; ok {
			data.Serials = append(data.Serials, serials[i].Seri)
			continue
		}

		data := &GenSerialRespData{
			ImportTicketID: serials[i].ImportTicket.ID,
			ItemSKU:        serials[i].Item.SKU,
			Serials:        []string{serials[i].Seri},
		}
		m[serials[i].Item.SKU] = data

		resp.Data = append(resp.Data, data)
	}
}
