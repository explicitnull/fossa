package httpserver

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type GetTicketsResponse struct {
	Message string
	Tickets []Ticket
}

type Ticket struct {
	ID          int
	Title       string
	Description string
}

func (s *Server) GetTickets(c *gin.Context) {
	ctx := c.Request.Context()

	tickets, err := s.ticketService.FetchTicketsFromJira(ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, c.Error(err))

		return
	}

	c.JSON(http.StatusOK, tickets)
}
