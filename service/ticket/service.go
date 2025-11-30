package ticket

import (
	"context"
	"fossa/pkg/logging"
	"fossa/service/template"
	"strings"

	"github.com/goccy/go-yaml"
	"github.com/pkg/errors"
)

const delimiterStart = "\nFor automation:"

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

// func (s *Service) FetchTicketsFromDB(ctx context.Context) ([]Ticket, error) {
// 	settings, err := s.repository.FetchTickets(ctx)
// 	if err != nil {
// 		return nil, errors.Wrap(err, "can't get tickets")
// 	}

// 	return settings, nil
// }

func (s *Service) FetchTicketsFromJira(ctx context.Context) ([]Ticket, error) {
	logger := logging.UnpackContext(ctx)

	tickets, err := s.jiraClient.FetchTickets(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "can't fetch tickets from Jira")
	}

	for _, t := range tickets {
		details, err := s.jiraClient.FetchTicketDetails(ctx, t.ID)
		if err != nil {
			return nil, errors.Wrap(err, "fetch ticket details")
		}

		vars, err := s.parseTicket(ctx, details)
		if err != nil {
			return nil, errors.Wrap(err, "parse ticket")
		}

		if len(vars) == 0 {
			logger.Debug("Skipping ticket %s as it does not contain template variables", t.ID)

			continue
		}

		logger.Debug("Parsed template variables for ticket %s: %+v", t.ID, vars)

		t.TemplateVariables = vars
	}

	return tickets, nil
}

func (s *Service) parseTicket(ctx context.Context, ticket *Ticket) (map[string]string, error) {
	des := strings.ToLower(ticket.Description)

	vars := make(map[string]string, 0)

	if !strings.Contains(des, delimiterStart) {
		return vars, nil
	}

	logger := logging.UnpackContext(ctx)
	logger.Debug("found automation directive in ticket %s!", ticket.ID)

	yml := strings.Split(des, delimiterStart)[1]

	err := yaml.Unmarshal([]byte(yml), &vars)
	if err != nil {
		return nil, errors.Wrap(err, "parse description")
	}

	return vars, nil
}
