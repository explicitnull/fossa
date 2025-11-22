package main

import (
	"context"
	"fossa/api/httpserver"
	"fossa/pkg/logging"
	"log"
	"os/signal"
	"syscall"

	"github.com/explicitnull/promcommon"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	cfg, err := NewConfig(defaultConfFileName)
	if err != nil {
		log.Fatalln(err)
	}

	logger, err := logging.NewLogger(&cfg.Logger, cfg.App.Name, promcommon.NewLoggerMetrics())
	if err != nil {
		log.Fatalln(err)
	}

	ctx = logging.PackContext(ctx, logger)

	// insert sqlite init

	// jira client init

	templatesRepository := templatesrepo.NewDB(sqliteConn)
	ticketsRepository := ticketsrepo.NewDB(sqliteConn)

	httpServer := httpserver.New()

	go func() {
		httpServer.Run()
	}()

	<-ctx.Done()

	log.Println("Initializing graceful shutdown")

	httpServer.Stop()

	log.Println("Shutdown complete")
}
