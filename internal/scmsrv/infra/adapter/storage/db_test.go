package storage_test

import (
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/thanhpp/scm/internal/scmsrv/infra/adapter/storage"
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

	return gorm.Open(
		postgres.Open(
			fmt.Sprintf(
				"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
				host, port, user, password, dbname, sslmode),
		),
		&gorm.Config{
			Logger: logger.New(
				log.New(os.Stdout, "\n", log.LstdFlags),
				logger.Config{
					LogLevel: logger.Info,
					Colorful: false,
				},
			),
			NamingStrategy: schema.NamingStrategy{
				SingularTable: true,
			},
		})
}

func TestDBTest(t *testing.T) {
	gDB, err := newPostgresDB()
	if err != nil {
		t.Error(err)

		return
	}

	db := storage.NewDB(gDB)

	if err := db.DropAll(); err != nil {
		t.Error(err)

		return
	}

	if err := db.AutoMigrate(); err != nil {
		t.Error(err)

		return
	}

	// if err := db.DropAll(); err != nil {
	// 	t.Error(err)

	// 	return
	// }
}
