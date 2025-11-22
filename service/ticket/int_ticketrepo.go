package ticket

import (
	"context"
)

type TicketRepository interface {
	FetchAll(ctx context.Context) ([]Ticket, error)
}
