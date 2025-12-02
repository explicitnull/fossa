package ticket

import (
	"context"
	"fossa/pkg/logging"
	"fossa/service/asset"
	"fossa/service/template"
	"strings"

	"github.com/goccy/go-yaml"
	"github.com/pkg/errors"
)

const delimiterStart = "\nfor automation:\n"

type Service struct {
	repository      TicketRepository
	templateService *template.Service
	assetService    *asset.Service
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
			continue
		}

		logger.Debug("Parsed template variables for ticket %s: %+v", t.ID, vars)

		t.TemplateVariables = vars
	}

	return tickets, nil
}

func (s *Service) parseTicket(ctx context.Context, ticket *Ticket) (map[string]interface{}, error) {
	des := strings.ToLower(ticket.Description)

	vars := make(map[string]interface{}, 0)

	if !strings.Contains(des, delimiterStart) {
		return vars, nil
	}

	yml := strings.Split(des, delimiterStart)[1]

	err := yaml.Unmarshal([]byte(yml), &vars)
	if err != nil {
		return nil, errors.Wrap(err, "parse description")
	}

	return vars, nil
}

// GetTickets is used by HTTP server
func (s *Service) GetTickets(ctx context.Context) ([]Ticket, error) {
	tickets, err := s.jiraClient.FetchTickets(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "can't fetch tickets from Jira")
	}

	res := make([]Ticket, 0, len(tickets))

	for _, t := range tickets {
		details, err := s.jiraClient.FetchTicketDetails(ctx, t.ID)
		if err != nil {
			return nil, errors.Wrap(err, "fetch ticket details")
		}

		vars, err := s.parseTicket(ctx, details)
		if err != nil {
			return nil, errors.Wrap(err, "parse ticket")
		}

		// Skip tickets without template variables
		if len(vars) == 0 {
			continue
		}

		t.TemplateVariables = vars
		res = append(res, t)

	}

	// fmt.Printf("######### Fetched %d tickets with template variables\n", len(res))

	return res, nil
}

// Used by HTTP server
func (s *Service) GetTicketByID(ctx context.Context, id string) (*Ticket, error) {
	tkt, err := s.jiraClient.FetchTicketDetails(ctx, id)
	if err != nil {
		return nil, errors.Wrap(err, "fetch ticket details")
	}

	vars, err := s.parseTicket(ctx, tkt)
	if err != nil {
		return nil, errors.Wrap(err, "parse ticket")
	}

	if len(vars) != 0 {
		tkt.TemplateVariables = vars
	}

	return tkt, nil
}
