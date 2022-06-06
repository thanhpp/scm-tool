package repo_test

import (
	"fmt"
	"testing"

	"github.com/thanhpp/scm/cmd/scmsrv/domain/repo"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

func newPostgresDB() (*gorm.DB, error) {
	const (
		host     = "localhost"
		port     = "5432"
		user     = "user"
		password = "password"
		dbname   = "scm"
		sslmode  = "disable"
	)
	const dsn = "host=localhost port=5432 user=user password=password dbname=scm sslmode=disable"

	return gorm.Open(
		postgres.Open(
			fmt.Sprintf(
				"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
				host, port, user, password, dbname, sslmode),
		),
		&gorm.Config{
			Logger: logger.Default.LogMode(logger.Info),
			NamingStrategy: schema.NamingStrategy{
				SingularTable: true,
			},
		})
}

func TestDBMigration(t *testing.T) {
	db, err := newPostgresDB()
	if err != nil {
		t.Error(err)

		return
	}

	var (
		tables []interface{} = []interface{}{
			repo.Item{}, repo.ItemType{}, repo.Serial{}, repo.Storage{},
		}
	)

	db.Migrator().DropTable(tables...)

	if err := db.AutoMigrate(tables...); err != nil {
		t.Error(err)

		return
	}

	db.Migrator().DropTable(tables...)
}
