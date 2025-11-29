package ticket

import (
	"context"
)

type JiraClient interface {
	FetchTicketsFromJira(ctx context.Context) ([]Ticket, error)
}
