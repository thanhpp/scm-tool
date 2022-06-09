package dto

import "time"

type CreateImportTicketReq struct {
	FromSupplierID int                            `json:"from_supplier_id"`
	ToStorageID    int                            `json:"to_storage_id"`
	SendTime       time.Time                      `json:"send_time"`
	Fee            float64                        `json:"fee"`
	Details        []CreateImportTicketReqDetails `json:"details"`
}

type CreateImportTicketReqDetails struct {
	ItemSKU     string  `json:"item_sku"`
	BuyQuantity int     `json:"buy_quantity"`
	BuyPrice    float64 `json:"buy_price"`
}
