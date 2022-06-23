package repo

type Supplier struct {
	ID    int    `gorm:"column:id; type:bigint; primaryKey; autoIncrement"`
	Name  string `gorm:"column:name; type:text"`
	Email string `gorm:"column:email; type:text"`
	Phone string `gorm:"column:phone; type:text"`
}
