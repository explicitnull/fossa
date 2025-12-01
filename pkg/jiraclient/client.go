package jiraclient

import (
	"context"
	"fmt"
	"fossa/service/ticket"

	jira "github.com/andygrunwald/go-jira"
	"github.com/pkg/errors"
)

const (
	maxSearchResults = 200
	allNetTicketsJQL = "project = NET AND status = \"In Progress\" ORDER BY created ASC"
)

type Config struct {
	Username string `env:"JIRA_USERNAME"`
	APIToken string `env:"JIRA_API_TOKEN"`
	URL      string `env:"JIRA_URL"`
}

type Client struct {
	client *jira.Client
}

func New(config Config) (*Client, error) {
	tp := jira.BasicAuthTransport{
		Username: config.Username,
		Password: config.APIToken,
	}

	client, err := jira.NewClient(tp.Client(), config.URL)
	if err != nil {
		return nil, errors.Wrap(err, "can't create Jira client: %v\n")
	}

	return &Client{client: client}, nil
}

func (c *Client) FetchTickets(ctx context.Context) ([]ticket.Ticket, error) {
	options := &jira.SearchOptionsV2{
		MaxResults: maxSearchResults,
		Fields:     []string{"summary", "status", "assignee", "description"},
	}

	issues, resp, err := c.client.Issue.SearchV2JQLWithContext(context.TODO(), allNetTicketsJQL, options)
	if err != nil {
		return nil, errors.Wrap(err, "can't search Jira issues")
	}
	if resp.StatusCode != 200 {
		return nil, errors.Errorf("non-200 response from Jira: %d", resp.StatusCode)
	}

	tt := make([]ticket.Ticket, 0, len(issues))

	for _, issue := range issues {
		fmt.Printf("%s:\n", issue.Key)

		t := ticket.Ticket{
			ID:    issue.Key,
			Title: issue.Fields.Summary,
		}
		tt = append(tt, t)
	}

	return tt, nil

}

func (c *Client) FetchTicketDetails(ctx context.Context, ticketID string) (*ticket.Ticket, error) {
	issue, resp, err := c.client.Issue.GetWithContext(ctx, ticketID, nil)
	if err != nil {
		return nil, errors.Wrap(err, "fetch issue")
	}
	if resp.StatusCode != 200 {
		return nil, errors.Errorf("non-200 response: %d", resp.StatusCode)
	}

	t := &ticket.Ticket{
		ID:          issue.Key,
		Title:       issue.Fields.Summary,
		Description: issue.Fields.Description,
	}

	return t, nil
}
