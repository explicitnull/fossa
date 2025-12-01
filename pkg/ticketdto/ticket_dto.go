package ticketdto

type GetTicketsResp struct {
	Message string   `json:"message"`
	Tickets []Ticket `json:"tickets"`
}

type GetTicketByIDResp struct {
	Message string `json:"message"`
	Ticket  Ticket `json:"ticket"`
}

type Ticket struct {
	ID          string  `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Assets      []Asset `json:"assets"`
}

type Asset struct {
	Step    string `json:"step"`
	Content string `json:"content"`
}
