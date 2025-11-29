package jiraclient

import (
	"context"
	"fmt"
	"fossa/service/ticket"

	jira "github.com/andygrunwald/go-jira"
	"github.com/pkg/errors"
)

const (
	maxSearchResults = 1000
	allNetTicketsJQL = "project = NET"
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

	fmt.Printf("Jira client created for user %s\n", config.Username)

	return &Client{client: client}, nil
}

func (c *Client) FetchTicketsFromJira(ctx context.Context) ([]ticket.Ticket, error) {
	fmt.Println("Fetching tickets from Jira")

	options := &jira.SearchOptionsV2{
		MaxResults: maxSearchResults,
		Fields:     []string{"summary", "status", "assignee"},
	}

	issues, resp, err := c.client.Issue.SearchV2JQLWithContext(context.TODO(), allNetTicketsJQL, options)
	if err != nil {
		fmt.Printf("Error searching JIRA client: %s\n", err)
		return nil, errors.Wrap(err, "can't search Jira issues")
	}
	if resp.StatusCode != 200 {
		fmt.Printf("Non-200 response: %d\n", resp.StatusCode)
		return nil, errors.Errorf("non-200 response from Jira: %d", resp.StatusCode)
	}

	tt := make([]ticket.Ticket, 0, len(issues))

	for _, issue := range issues {
		fmt.Printf("%s:\n", issue.Key)
		fmt.Printf(" %s:\n", issue.Fields.Summary)

		t := ticket.Ticket{
			ID:    issue.Key,
			Title: issue.Fields.Summary,
		}
		tt = append(tt, t)
	}

	return tt, nil

}
