package repo

import "time"

type Item struct {
	SKU        string  `gorm:"column:sku;type:text; primaryKey"`
	Name       string  `gorm:"column:name;type:text"`
	Desc       string  `gorm:"column:desc;type:text"`
	SellPrice  float64 `gorm:"column:sell_price;type:float(8)"`
	ItemTypeID int     `gorm:"column:item_type_id;type:int"`

	Images   []ItemImage `gorm:"foreignKey:ItemSKU; references:SKU"`
	ItemType ItemType    `gorm:"foreignKey:ItemTypeID; references:ID"`
	Serials  []Serial    `gorm:"foreignKey:ItemSKU;references:SKU"`

	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt time.Time `gorm:"column:updated_at;autoUpdateTime:milli"`
}

type ItemImage struct {
	ItemSKU string `gorm:"column:item_sku;type:text"`
	Image   string `gorm:"column:image;type:text"`
}

type ItemType struct {
	ID   int    `gorm:"column:id; type:int; primaryKey; autoIncrement"`
	Name string `gorm:"column:name;type:text; unique"`
	Desc string `gorm:"column:desc;type:text"`

	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt time.Time `gorm:"column:updated_at;autoUpdateTime:milli"`
}
