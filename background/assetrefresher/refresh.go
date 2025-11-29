package assetrefresher

import (
	"context"
	"fossa/pkg/logging"
	"fossa/service/ticket"
	"time"
)

const updateInterval = 1 * time.Minute

type Refresher struct {
	tkr           *time.Ticker
	ticketService *ticket.Service
	logger        *logging.Logger
}

func New(ticketService *ticket.Service, logger *logging.Logger) *Refresher {
	return &Refresher{
		tkr:           time.NewTicker(updateInterval),
		ticketService: ticketService,
		logger:        logger,
	}
}

func (r *Refresher) Run(ctx context.Context) {
	for range r.tkr.C {
		err := r.ticketService.GenerateTexts(ctx)
		if err != nil {
			r.logger.Error("Error generating texts: %v", err)
		}
	}
}

func (r *Refresher) Stop() {
	r.tkr.Stop()
}
