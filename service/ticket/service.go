package ticket

import (
	"context"
	"encoding/json"
	"fmt"
	"fossa/pkg/history"
	"fossa/service/template"

	"github.com/pkg/errors"
)

type Service struct {
	repository      TicketRepository
	templateService *template.Service
	jiraClient      JiraClient
}

func NewService(
	repository TicketRepository,
	templateService *template.Service,
	jiraClient JiraClient,
) *Service {
	return &Service{
		repository:      repository,
		templateService: templateService,
		jiraClient:      jiraClient,
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

func (s *Service) GenerateTexts(ctx context.Context) error {
	tickets, err := s.jiraClient.FetchTicketsFromJira(ctx)
	if err != nil {
		return errors.Wrap(err, "can't fetch tickets from Jira")
	}

	// allAssets := make(map[string]ticketAssets)

	for _, t := range tickets {
		fmt.Printf("Processing ticket: %s\n", t.ID)
	}
	return nil
}
