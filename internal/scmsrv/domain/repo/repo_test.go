package repo_test

import (
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/thanhpp/scm/internal/scmsrv/domain/repo"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

func newPostgresDB() (*gorm.DB, error) {
	const (
		host     = "localhost"
		port     = "5432"
		user     = "scmuser"
		password = "scmpassword"
		dbname   = "scmdb"
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

func TestDBMigration(t *testing.T) {
	db, err := newPostgresDB()
	if err != nil {
		t.Error(err)

		return
	}

	var (
		tables []interface{} = []interface{}{
			repo.ItemType{},
		}
	)

	_ = db.Migrator().DropTable(tables...)

	if err := db.AutoMigrate(tables...); err != nil {
		t.Error(err)

		return
	}

	_ = db.Migrator().DropTable(tables...)
}
