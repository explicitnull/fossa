package ticket

import (
	"context"
	"encoding/json"
	"fossa/pkg/history"
	"fossa/service/template"

	"github.com/pkg/errors"
)

type Service struct {
	repository      TicketRepository
	templateService *template.Service
}

func NewService(
	repository TicketRepository,
	templateService *template.Service,
	// jiraClient jira.Client,
) *Service {
	return &Service{
		repository:      repository,
		templateService: templateService,
	}
}

type Ticket struct {
	ID          string
	Title       string
	Priority    string
	Description json.RawMessage
	history.HistoryData
}

func (s *Service) FetchTicketsFromDB(ctx context.Context) ([]Ticket, error) {
	/*
		- fetch tickets with filter
		- calculate hash on json/yaml
		- compare hash with cached value, proceed if not changed
		-
	*/
	settings, err := s.repository.FetchTickets(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "can't get tickets")
	}

	return settings, nil
}
