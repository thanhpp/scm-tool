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
	"github.com/thanhpp/scm/internal/nftsrv/nftcfg"
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
		logger.Warnw("load env file error", "file", constx.DefaultENVFile, "err", err)
	}

	cfg, err := nftcfg.NewNFTServiceConfig(constx.DefaultConfigFile)
	if err != nil {
		logger.Fatalw("load config error", "file", constx.DefaultConfigFile, "err", err)
	}

	if err := logger.SetLog(cfg.Logger); err != nil {
		logger.Fatalw("set log error", "config", cfg.Logger, "err", err)
	}

	ipfsClient, err := ipfsclient.NewIPFSClient(cfg.InfuraConfig.ProjectID, cfg.InfuraConfig.ProjectSecret)
	if err != nil {
		logger.Fatalw("create ipfs client error", "err", err)
	}

	mainCtx, cancel := context.WithCancel(context.Background())
	defer cancel()

	ethClient, err := ethclient.DialContext(mainCtx, cfg.NodeAPIURL)
	if err != nil {
		logger.Fatalw("create rinkeby eth client error", "err", err)
	}

	nftMinter, err := nftminter.NewNFTMinter(ethClient, cfg.PrivateKey, smartcontracts.ContractAddress)
	if err != nil {
		logger.Fatalw("create nft minter error", "err", err)
	}

	gdb, err := newGormDB(cfg.Database.DSN())
	if err != nil {
		logger.Fatalw("create gorm db error", "err", err)
	}

	storage := storage.NewDB(gdb)

	nftApp := app.NewApp(mainCtx, ipfsClient, nftMinter, storage.SeriNFTRepo())
	logger.Debugw("start nft app", "app", nftApp)
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
