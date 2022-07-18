package main

import (
	"log"
	"os"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/joho/godotenv"
	"github.com/thanhpp/scm/internal/nftsrv/app"
	"github.com/thanhpp/scm/internal/nftsrv/infra/adapter/ipfsclient"
	"github.com/thanhpp/scm/internal/nftsrv/infra/adapter/nftminter"
	"github.com/thanhpp/scm/internal/nftsrv/infra/adapter/storage"
	"github.com/thanhpp/scm/internal/nftsrv/infra/port/httpsv"
	"github.com/thanhpp/scm/internal/nftsrv/nftcfg"
	"github.com/thanhpp/scm/pkg/booting"
	"github.com/thanhpp/scm/pkg/constx"
	"github.com/thanhpp/scm/pkg/logger"
	"github.com/thanhpp/scm/pkg/smartcontracts"
	"golang.org/x/net/context"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	gormlogger "gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

func main() {
	if err := godotenv.Load(constx.DefaultENVFile); err != nil {
		log.Println("load env file error", "file", constx.DefaultENVFile, "err", err)
		return
	}

	cfg, err := nftcfg.NewNFTServiceConfig(constx.DefaultConfigFile)
	if err != nil {
		log.Println("load config error", "file", constx.DefaultConfigFile, "err", err)
		return
	}
	log.Println("read config OK")

	if err := logger.SetLog(cfg.Logger); err != nil {
		logger.Fatalw("set log error", "config", cfg.Logger, "err", err)
	}
	logger.Info("set log OK")

	ipfsClient, err := ipfsclient.NewIPFSClient(cfg.InfuraConfig.ProjectID, cfg.InfuraConfig.ProjectSecret)
	if err != nil {
		logger.Fatalw("create ipfs client error", "err", err)
	}
	logger.Info("ipfs client OK")

	mainCtx, cancel := context.WithCancel(context.Background())
	defer cancel()

	ethClient, err := ethclient.DialContext(mainCtx, cfg.NodeAPIURL)
	if err != nil {
		logger.Fatalw("create rinkeby eth client error", "err", err)
	}
	logger.Info("rinkeby eth client OK")

	nftMinter, err := nftminter.NewNFTMinter(ethClient, cfg.PrivateKey, smartcontracts.ContractAddress)
	if err != nil {
		logger.Fatalw("create nft minter error", "err", err)
	}
	logger.Info("nft minter OK")

	gdb, err := newGormDB(cfg.Database.DSN())
	if err != nil {
		logger.Fatalw("create gorm db error", "err", err)
	}
	logger.Info("gorm db OK")

	storage := storage.NewDB(gdb)

	nftApp := app.NewApp(mainCtx, ipfsClient, nftMinter, storage.SeriNFTRepo())
	logger.Debugw("start nft app", "app", nftApp)

	server := httpsv.NewHTTPServer(cfg.HTTPServer, nftApp)

	daemonMan := booting.NewDaemonManeger(mainCtx)
	daemonMan.Start(server.Daemon())
	booting.WaitSignals(mainCtx)
	daemonMan.Stop()
}

func newGormDB(dsn string) (*gorm.DB, error) {
	return gorm.Open(
		postgres.Open(dsn),
		&gorm.Config{
			Logger: gormlogger.New(
				log.New(os.Stdout, "\n", log.LstdFlags),
				gormlogger.Config{
					LogLevel: gormlogger.Error,
					Colorful: false,
				},
			),
			NamingStrategy: schema.NamingStrategy{
				SingularTable: true,
			},
		})
}
