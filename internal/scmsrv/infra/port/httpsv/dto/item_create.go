package dto

import (
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
	SKU        string `json:"sku" form:"sku"`
	Name       string `json:"name" form:"name"`
	Desc       string `json:"desc" form:"desc"`
	ItemTypeID int    `json:"item_type_id" form:"item_type_id"`
}
type ItemInfoResp struct {
	ginutil.RespTemplateError
	Data ItemInfoRespData `json:"data"`
}

func (resp *ItemInfoResp) SetData(item *entity.Item) {
	resp.Data = ItemInfoRespData{
		SKU:        item.SKU,
		Name:       item.Name,
		Desc:       item.Desc,
		ItemTypeID: item.Type.ID,
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

type ItemTypeInfoResp struct {
	ginutil.RespTemplateError
	Data ItemTypeInfoRespData `json:"data"`
}

func (resp *ItemTypeInfoResp) SetData(itemType *entity.ItemType) {
	resp.Data = ItemTypeInfoRespData{
		ID:   itemType.ID,
		Name: itemType.Name,
		Desc: itemType.Desc,
	}
}
