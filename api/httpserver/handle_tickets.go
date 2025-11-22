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
	c.JSON(http.StatusOK, nil)
}
