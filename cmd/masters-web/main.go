package main

import (
	"context"
	"fmt"
	"github.com/nkralles/masters-web/internal/cfg"
	"github.com/nkralles/masters-web/internal/logger"
	"github.com/nkralles/masters-web/internal/persistence"
	"github.com/nkralles/masters-web/internal/persistence/pgdriver"
	"github.com/nkralles/masters-web/internal/routes"
)

func main() {

	var driver persistence.MastersStorage
	var err error
	driver, err = pgdriver.Connect(cfg.Config.DatabaseURL)
	if err != nil {
		logger.Internal.Fatalf("storage driver 	init failed: %v", err)
	}
	persistence.SetDefaultDriver(driver)

	logger.Internal.Debugf("%+v\n", cfg.Config)
	addr := fmt.Sprintf("%s:%d", cfg.Config.ListenAddress, cfg.Config.Port)
	logger.Internal.Infof("Starting server to listen on %s", addr)
	err = routes.ListenAndServe(context.Background(), addr)
	if err != nil {
		logger.Internal.Fatal(err)
	}
}
