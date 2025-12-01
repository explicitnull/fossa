package main

import (
	"context"
	"fossa/api/httpserver"
	"fossa/background/assetrefresher"
	"fossa/pkg/jiraclient"
	"fossa/pkg/logging"
	"fossa/pkg/sqlite"
	"fossa/service/asset"
	"fossa/service/template"
	"fossa/service/ticket"
	"log"
	"os/signal"
	"syscall"

	"fossa/repository/templaterepo"
	"fossa/repository/ticketrepo"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	cfg, err := NewConfig(defaultConfFileName)
	if err != nil {
		log.Fatalln(err)
	}

	logger, err := logging.NewLogger(&cfg.Logger, cfg.App.Name)
	if err != nil {
		log.Fatalln(err)
	}

	ctx = logging.PackContext(ctx, logger)

	sqliteConn, err := sqlite.NewDB()
	if err != nil {
		logger.Fatal("Can't initialize SQLite")
	}

	jiraClient, err := jiraclient.New(cfg.Jira)
	if err != nil {
		logger.Fatal("Can't initialize Jira")
	}

	templatesRepository := templaterepo.NewSQLite(sqliteConn)
	ticketsRepository := ticketrepo.NewSQLite(sqliteConn)

	templatesService := template.NewService(templatesRepository)
	ticketsService := ticket.NewService(ticketsRepository, templatesService, jiraClient)
	assetService := asset.NewService(nil, templatesService)

	httpServer := httpserver.New(cfg.HTTPServer, logger, ticketsService, assetService)

	go func() {
		httpServer.Run()
	}()

	backgroundAssetRefresher := assetrefresher.New(
		ticketsService,
		assetService,
		logger,
	)

	go func() {
		backgroundAssetRefresher.Run(ctx)
	}()

	<-ctx.Done()

	logger.Warn("System signal received, initializing graceful shutdown")

	httpServer.Stop()
	// backgroundassetrefresher.Stop()

	logger.Warn("Shutdown complete")
}
