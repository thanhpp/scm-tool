package storage

import (
	"github.com/thanhpp/scm/internal/nftsrv/domain/repo"
	"github.com/thanhpp/scm/pkg/logger"
	"gorm.io/gorm"
)

var collection = []interface{}{repo.SeriNFT{}}

type DB struct {
	gdb *gorm.DB
}

func NewDB(gormDB *gorm.DB) *DB {
	db := &DB{
		gdb: gormDB,
	}

	if err := db.AutoMigrate(); err != nil {
		logger.Fatalf("migrate db err %v", err)
	}

	return db
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
