package repo

import (
	"time"

	"github.com/google/uuid"
)

type Storage struct {
	ID       uuid.UUID `gorm:"column:id;type:uuid;primaryKey"`
	Name     string    `gorm:"column:name;type:text"`
	Location string    `gorm:"column:name;type:text"`

	Serials []Serial `gorm:"foreignKey:StorageID;references:ID"`

	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt time.Time `gorm:"column:updated_at;autoUpdateTime:milli"`
}
