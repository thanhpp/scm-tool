package repo

import (
	"time"

	"github.com/thanhpp/scm/pkg/enum"
)

type Serial struct {
	Seri           string            `gorm:"column:seri; type:text; primaryKey"`
	Status         enum.SerialStatus `gorm:"column:status"`
	ItemSKU        string            `gorm:"column:item_sku; type:text; primaryKey"`
	ImportTicketID int               `gorm:"column:import_ticket_id; type:int"`
	TokenID        int               `gorm:"column:token_id; type:bigint; default:0"`

	Item         Item         `gorm:"foreignKey:ItemSKU;references:SKU"`
	ImportTicket ImportTicket `gorm:"foreignKey:ImportTicketID;references:ID"`

	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt time.Time `gorm:"column:updated_at;autoUpdateTime:milli"`
}
