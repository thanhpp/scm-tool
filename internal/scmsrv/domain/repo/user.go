package repo

type User struct {
	ID           int    `gorm:"column:id; type:int; primaryKey; autoIncrement"`
	Name         string `gorm:"column:name; type:text"`
	Username     string `gorm:"column:username; type:text"`
	HashPassword string `gorm:"column:hash_password; type:text"`
}
