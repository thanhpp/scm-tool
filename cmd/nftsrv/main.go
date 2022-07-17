package main

import (
	"github.com/joho/godotenv"
	"github.com/thanhpp/scm/internal/nftsrv/nftcfg"
	"github.com/thanhpp/scm/pkg/constx"
	"github.com/thanhpp/scm/pkg/logger"
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
}
