package main

import (
	"context"
	"fmt"
	"log"

	"github.com/thanhpp/scm/internal/scmsrv/infra/port/httpsv"
	"github.com/thanhpp/scm/internal/scmsrv/scmcfg"
	"github.com/thanhpp/scm/pkg/booting"
	"github.com/thanhpp/scm/pkg/configx"
	"github.com/thanhpp/scm/pkg/logger"
)

func main() {
	mainCfg := new(scmcfg.MainConfig)

	if err := configx.ReadConfigFromFile("config.yml", mainCfg); err != nil {
		panic(err)
	}

	if err := logger.SetLog(mainCfg.Logger); err != nil {
		log.Fatal("set log err", err)
	}

	httpServer := httpsv.NewHTTPServer(mainCfg.HTTPServer)

	mainCtx := context.Background()
	daemonMan := booting.NewDaemonManeger(mainCtx)
	daemonMan.Start(httpServer.Daemon())

	booting.WaitSignals(mainCtx)

	daemonMan.Stop()

	fmt.Println(mainCfg)
}
