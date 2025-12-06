package ticketdto

import "fossa/pkg/dto/assetdto"

type GetTicketsResp struct {
	Message string   `json:"message"`
	Tickets []Ticket `json:"tickets"`
}

type GetTicketByIDResp struct {
	Message string `json:"message"`
	Ticket  Ticket `json:"ticket"`
}

type Ticket struct {
	ID          string           `json:"id"`
	Title       string           `json:"title"`
	Description string           `json:"description,omitempty"`
	Assignee    string           `json:"assignee,omitempty"`
	Assets      []assetdto.Asset `json:"assets,omitempty"`
}
