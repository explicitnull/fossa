package ticket

import (
	"context"
)

type JiraClient interface {
	FetchTickets(ctx context.Context) ([]Ticket, error)
	FetchTicketDetails(ctx context.Context, ticketID string) (*Ticket, error)
}
