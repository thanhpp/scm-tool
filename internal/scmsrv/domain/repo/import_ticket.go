package repo

type ImportTicket struct {
	ID int `gorm:"column:id;type:int;primaryKey;autoIncrement"`

	ToStorage Storage
}
