package dto

import (
	"fmt"
	"path/filepath"

	"github.com/thanhpp/scm/internal/scmsrv/domain/entity"
	"github.com/thanhpp/scm/pkg/ginutil"
)

type CreateItemReq struct {
	SKU        string `json:"sku" form:"sku"`
	Name       string `json:"name" form:"name"`
	Desc       string `json:"desc" form:"desc"`
	ItemTypeID int    `json:"item_type_id" form:"item_type_id"`
}

type ItemInfoRespData struct {
	SKU        string   `json:"sku" form:"sku"`
	Name       string   `json:"name" form:"name"`
	Desc       string   `json:"desc" form:"desc"`
	ItemTypeID int      `json:"item_type_id" form:"item_type_id"`
	Images     []string `json:"images"`
}

func (d *ItemInfoRespData) set(item *entity.Item) {
	d.SKU = item.SKU
	d.Name = item.Name
	d.Desc = item.Desc
	d.ItemTypeID = item.Type.ID
	d.Images = make([]string, len(item.Images))

	for i := range d.Images {
		d.Images[i] = buildFileURL(item.Images[i])
	}
}

func buildFileURL(filename string) string {
	return fmt.Sprintf("%s/files/%s", "https://scm-tool.thanhpp.ninja", filepath.Base(filename))
}

type ItemInfoResp struct {
	ginutil.RespTemplateError
	Data ItemInfoRespData `json:"data"`
}

func (resp *ItemInfoResp) SetData(item *entity.Item) {
	resp.Data.set(item)
}

type GetListItemInfoResp struct {
	ginutil.RespTemplateError
	Data []ItemInfoRespData `json:"data"`
}

func (resp *GetListItemInfoResp) SetData(items []*entity.Item) {
	resp.Data = make([]ItemInfoRespData, len(items))

	for i := range items {
		resp.Data[i].set(items[i])
	}
}

type CreateItemTypeReq struct {
	Name string `json:"name"`
	Desc string `json:"desc"`
}

type ItemTypeInfoRespData struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Desc string `json:"desc"`
}

func (d *ItemTypeInfoRespData) set(itemType *entity.ItemType) {
	d.ID = itemType.ID
	d.Name = itemType.Name
	d.Desc = itemType.Desc
}

type ItemTypeInfoResp struct {
	ginutil.RespTemplateError
	Data ItemTypeInfoRespData `json:"data"`
}

func (resp *ItemTypeInfoResp) SetData(itemType *entity.ItemType) {
	resp.Data.set(itemType)
}

type GetAllItemTypeResp struct {
	ginutil.RespTemplateError
	Data []ItemTypeInfoRespData `json:"data"`
}

func (resp *GetAllItemTypeResp) SetData(itemTypes []*entity.ItemType) {
	resp.Data = make([]ItemTypeInfoRespData, len(itemTypes))

	for i := range itemTypes {
		resp.Data[i].set(itemTypes[i])
	}
}

type ReqUpdateItemType struct {
	Name string `json:"name"`
	Desc string `json:"desc"`
}

type ReqUpdateItem struct {
	Name         string   `form:"name"`
	Desc         string   `form:"desc"`
	ItemTypeID   int      `form:"item_type_id"`
	SellPrice    float64  `form:"sell_price"`
	DeleteImages []string `form:"delete_images"`
}
