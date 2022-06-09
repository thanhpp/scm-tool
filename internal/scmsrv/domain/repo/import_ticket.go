package repo

import (
	"time"

	"github.com/thanhpp/scm/pkg/enum"
)

type ImportTicket struct {
	ID             int                     `gorm:"column:id;type:int;primaryKey;autoIncrement"`
	FromSupplierID int                     `gorm:"column:from_supplier_id; type:int"`
	ToStorageID    int                     `gorm:"column:to_storage_id; type:int"`
	Status         enum.ImportTicketStatus `gorm:"column:status"`
	SendTime       time.Time               `gorm:"column:send_time"`
	ReceiveTime    time.Time               `gorm:"column:receive_time"`
	Fee            float64                 `gorm:"column:fee; type:float8"`

	FromSupplier  Supplier                   `gorm:"foreignKey:FromSupplierID; references:ID"`
	ToStorage     Storage                    `gorm:"foreignKey:ToStorageID; references:ID"`
	ProductImages []ImportTicketProductImage `gorm:"foreignKey:ImportTicketID; references:ID"`
	BillImages    []ImportTicketBillImage    `gorm:"foreignKey:ImportTicketID; references:ID"`
	Details       []ImportTicketDetails      `gorm:"foreignKey:ImportTicketID; references:ID"`

	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt time.Time `gorm:"column:updated_at;autoUpdateTime:milli"`
}

type ImportTicketProductImage struct {
	ImportTicketID   int       `gorm:"column:import_ticket_id; type:int"`
	ProductImagePath string    `gorm:"column:product_image_path; type:text"`
	CreatedAt        time.Time `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt        time.Time `gorm:"column:updated_at;autoUpdateTime:milli"`
}

type ImportTicketBillImage struct {
	ImportTicketID int       `gorm:"column:import_ticket_id;type:int"`
	BillImagePath  string    `gorm:"column:bill_image_path;type:text"`
	CreatedAt      time.Time `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt      time.Time `gorm:"column:updated_at;autoUpdateTime:milli"`
}

type ImportTicketDetails struct {
	ImportTicketID  int     `gorm:"column:import_ticket_id;type:int"`
	ItemSKU         string  `gorm:"column:item_sku;type:text"`
	BuyQuantity     int     `gorm:"column:buy_quantity;type:int"`
	ReceiveQuantity int     `gorm:"column:receive_quantity;type:int"`
	BuyPrice        float64 `gorm:"buy_price;type:float8"`
}
