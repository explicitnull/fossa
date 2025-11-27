package ticket

import (
	"context"
)

type TicketRepository interface {
	FetchTickets(ctx context.Context) ([]Ticket, error)
}
