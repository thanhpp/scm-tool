package repo

import (
	"time"
)

type Serial struct {
	Seri           string `gorm:"column:seri;type:text;primaryKey"`
	ItemSKU        string `gorm:"column:item_sku;type:text;primaryKey"`
	StorageID      int    `gorm:"column:storage_id;type:int"`
	ImportTicketID int    `gorm:"column:import_ticket_id;type:int"`

	Storage      Storage      `gorm:"foreignKey:StorageID;references:ID"`
	ImportTicket ImportTicket `gorm:"foreignKey:ImportTicketID;references:ID"`

	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt time.Time `gorm:"column:updated_at;autoUpdateTime:milli"`
}
