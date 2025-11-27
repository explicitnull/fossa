package main

import (
	"context"
	"fossa/api/httpserver"
	"fossa/pkg/logging"
	"fossa/pkg/sqlite"
	"fossa/service/template"
	"fossa/service/ticket"
	"log"
	"os/signal"
	"syscall"

	"fossa/repository/templaterepo"
	"fossa/repository/ticketrepo"
	// go-jira "github.com/andygrunwald/go-jira/v1"
)

const jiraURL = "https://cenic.atlassion.com/jira/"

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

	// jiraClient, err := jira.NewClient(nil, "https://issues.apache.org/jira/")
	// if err != nil {
	// 	logger.Fatal("Can't initialize Jira")
	// }

	// _ = jiraClient

	templatesRepository := templaterepo.NewSQLite(sqliteConn)
	ticketsRepository := ticketrepo.NewSQLite(sqliteConn)

	templatesService := template.NewService(templatesRepository)
	ticketsService := ticket.NewService(ticketsRepository, templatesService)

	httpServer := httpserver.New(cfg.HTTPServer, ticketsService)

	go func() {
		httpServer.Run()
	}()

	// backgroundJiraFetcher := jiraFetcher.New(
	// 	logger,
	// 	jiraFetcher,
	// )

	// go func() {
	// 	backgroundJiraFetcher.Run(ctx)
	// }()

	<-ctx.Done()

	logger.Warn("Initializing graceful shutdown")

	httpServer.Stop()
	// backgroundJiraFetcher.Stop()

	logger.Warn("Shutdown complete")
}
