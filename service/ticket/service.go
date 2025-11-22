package ticket

import (
	"context"
	"encoding/json"

	"github.com/pkg/errors"
)

type Service struct {
	repository TicketRepository
}

func NewService(
	repository TicketRepository,
	jiraClient JiraClient,
) *Service {
	return &Service{
		repository: repository,
	}
}

type Ticket struct {
	ID          string
	Title       string
	Priority    string
	Reported    string
	Description json.RawMessage
}

func (s *Service) GetTickets(ctx context.Context) ([]Ticket, error) {
	/*
		- fetch tickets with filter
		- calculate hash on json/yaml
		- compare hash with cached value, proceed if not changed
		-
	*/
	settings, err := s.repository.FetchAll(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "can't get tickets")
	}

	return settings, nil
}
