package ticketdto

type TicketsResp struct {
	Message string   `json:"message"`
	Tickets []Ticket `json:"tickets"`
}

type Ticket struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}
