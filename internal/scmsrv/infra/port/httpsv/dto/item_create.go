package dto

type CreateItemReq struct {
	SKU        string `json:"sku" form:"sku"`
	Name       string `json:"name" form:"name"`
	Desc       string `json:"desc" form:"desc"`
	ItemTypeID int    `json:"item_type_id" form:"item_type_id"`
}
