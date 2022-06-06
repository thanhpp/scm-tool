package repo

import (
	"time"

	"github.com/google/uuid"
)

type Serial struct {
	Seri      string    `gorm:"column:seri;type:text;primaryKey"`
	ItemSKU   string    `gorm:"column:item_sku;type:text;primaryKey"`
	StorageID uuid.UUID `gorm:"column:storage_id;type:uuid"`

	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt time.Time `gorm:"column:updated_at;autoUpdateTime:milli"`
}
