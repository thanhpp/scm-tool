package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/thanhpp/scm/internal/scmsrv/app"
	"github.com/thanhpp/scm/internal/scmsrv/domain/entity"
	"github.com/thanhpp/scm/internal/scmsrv/infra/adapter/nftsvclient"
	"github.com/thanhpp/scm/internal/scmsrv/infra/adapter/storage"
	"github.com/thanhpp/scm/internal/scmsrv/infra/port/httpsv"
	"github.com/thanhpp/scm/internal/scmsrv/scmcfg"
	"github.com/thanhpp/scm/pkg/booting"
	"github.com/thanhpp/scm/pkg/configx"
	"github.com/thanhpp/scm/pkg/constx"
	"github.com/thanhpp/scm/pkg/fileutil"
	"github.com/thanhpp/scm/pkg/logger"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	gormlogger "gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

// func init() {
// 	// set timezone
// 	tz, err := time.LoadLocation("Asia/Bangkok")
// 	if err != nil {
// 		panic(err)
// 	}

// 	time.Local = tz
// }

func main() {
	mainCfg := new(scmcfg.MainConfig)

	_ = os.MkdirAll(constx.SaveFilePaths, os.ModePerm)

	if err := configx.ReadConfigFromFile("config.yml", mainCfg); err != nil {
		panic(err)
	}

	mainCfg.Database.OverideWithEnv()

	// setup
	if err := logger.SetLog(mainCfg.Logger); err != nil {
		log.Fatal("set log err", err)
	}

	logger.Debugf("dsn %s", mainCfg.Database.DSN())
	gdb, err := newGormDB(mainCfg.Database.DSN())
	if err != nil {
		panic(err)
	}

	db := storage.NewDB(gdb)
	nftServiceClient := nftsvclient.New(constx.DefaultNFTServiceURL)

	scmApp := app.New(
		entity.NewFactory(),
		db.ItemDB(), db.SupplierDB(), db.StorageDB(),
		db.ImportTicketDB(), db.SerialDB(), db.UserDB(),
		fileutil.NewFileUtil(), nftServiceClient,
	)

	httpServer := httpsv.NewHTTPServer(mainCfg.HTTPServer, &scmApp)

	autoMintAndUpdateSerialDaemon := scmApp.AutoMintAndUpdateSerial(db.SerialDB(), nftServiceClient)

	// start
	mainCtx := context.Background()
	daemonMan := booting.NewDaemonManeger(mainCtx)
	daemonMan.Start(httpServer.Daemon(), autoMintAndUpdateSerialDaemon)

	booting.WaitSignals(mainCtx)

	daemonMan.Stop()

	fmt.Println(mainCfg)
}

func newGormDB(dsn string) (*gorm.DB, error) {
	return gorm.Open(
		postgres.Open(dsn),
		&gorm.Config{
			Logger: gormlogger.New(
				log.New(os.Stdout, "\n", log.LstdFlags),
				gormlogger.Config{
					LogLevel: gormlogger.Info,
					Colorful: false,
				},
			),
			NamingStrategy: schema.NamingStrategy{
				SingularTable: true,
			},
		})
}
