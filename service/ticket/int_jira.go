package ticket

import (
	"context"
)

type JiraClient interface {
	FetchTickets(ctx context.Context, query string) ([]Ticket, error)
	FetchTicketDetails(ctx context.Context, ticketID string) (*Ticket, error)
}
