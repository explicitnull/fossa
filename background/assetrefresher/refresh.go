package assetrefresher

import (
	"context"
	"fossa/pkg/logging"
	"fossa/service/asset"
	"fossa/service/ticket"
	"time"
)

const updateInterval = 10 * time.Minute

type Refresher struct {
	tkr           *time.Ticker
	ticketService *ticket.Service
	assetService  *asset.Service
	logger        *logging.Logger
}

func New(ticketService *ticket.Service, assetService *asset.Service, logger *logging.Logger) *Refresher {
	return &Refresher{
		tkr:           time.NewTicker(updateInterval),
		ticketService: ticketService,
		assetService:  assetService,
		logger:        logger,
	}
}

func (r *Refresher) Run(ctx context.Context) {
	r.logger.Debug("Asset refresher started")

	for range r.tkr.C {
		err := r.generateAssets(ctx)
		if err != nil {
			r.logger.Error("generate texts: %v", err)
		}
	}
}

func (r *Refresher) Stop() {
	r.tkr.Stop()
}

func (r *Refresher) generateAssets(ctx context.Context) error {
	r.logger.Debug("Asset refresher triggered")

	/*
		TODO:
		- calculate hash on json/yaml in Jira
		- compare hash with cached value, proceed if hash changed
	*/

	tickets, err := r.ticketService.FetchTicketsFromJira(ctx)
	if err != nil {
		return err
	}

	for _, t := range tickets {
		if t.TemplateVariables == nil {
			continue
		}

		_, err := r.assetService.GenerateAssetsForTicket(ctx, t.TemplateVariables)
		if err != nil {
			r.logger.Error("generate assets %s: %v", t.ID, err)
			continue
		}
	}

	return nil
}
