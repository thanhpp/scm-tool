package storage

import (
	"github.com/thanhpp/scm/internal/scmsrv/domain/repo"
	"gorm.io/gorm"
)

var (
	collection []interface{} = []interface{}{
		repo.Item{}, repo.ItemType{}, repo.Storage{}, repo.Serial{},
		repo.Supplier{}, repo.ImportTicket{},
		repo.ImportTicketBillImage{}, repo.ImportTicketProductImage{}, repo.ImportTicketDetails{},
	}
)

type DB struct {
	gdb *gorm.DB
}

func NewDB(gormDB *gorm.DB) *DB {
	return &DB{
		gdb: gormDB,
	}
}

func (d DB) AutoMigrate() error {
	return d.gdb.AutoMigrate(
		collection...,
	)
}

func (d DB) DropAll() error {
	return d.gdb.Migrator().DropTable(
		collection...,
	)
}
