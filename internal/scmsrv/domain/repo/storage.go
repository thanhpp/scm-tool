package repo

import (
	"time"
)

type Storage struct {
	ID       int    `gorm:"column:id;type:int;primaryKey;autoIncrement"`
	Name     string `gorm:"column:name;type:text"`
	Location string `gorm:"column:name;type:text"`

	Serials []Serial `gorm:"foreignKey:StorageID;references:ID"`

	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt time.Time `gorm:"column:updated_at;autoUpdateTime:milli"`
}