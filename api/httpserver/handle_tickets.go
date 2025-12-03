package httpserver

import (
	"fossa/pkg/ticketdto"
	"fossa/service/asset"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *Server) GetTickets(c *gin.Context) {
	ctx := c.Request.Context()

	tickets, err := s.ticketService.GetTickets(ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, c.Error(err))

		return
	}

	allTicketsDTO := make([]ticketdto.Ticket, 0, len(tickets))

	for _, tkt := range tickets {
		ticketDTO := ticketdto.Ticket{
			ID:       tkt.ID,
			Title:    tkt.Title,
			Assignee: tkt.Assignee,
		}

		allTicketsDTO = append(allTicketsDTO, ticketDTO)
	}

	result := ticketdto.GetTicketsResp{
		Message: "",
		Tickets: allTicketsDTO,
	}

	c.JSON(http.StatusOK, result)
}

func (s *Server) GetTicketByID(c *gin.Context) {
	ctx := c.Request.Context()

	ticketID := c.Param("id")

	tkt, err := s.ticketService.GetTicketByID(ctx, ticketID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, c.Error(err))

		return
	}

	assets, err := s.assetService.GenerateAssetsForTicket(ctx, tkt.TemplateVariables)
	if err != nil && err != asset.ErrJobTypeNotFound {
		c.JSON(http.StatusInternalServerError, c.Error(err))
		return
	}

	// sending response
	assetsDTO := make([]ticketdto.Asset, 0, len(assets))
	for step, asset := range assets {
		assetsDTO = append(assetsDTO, ticketdto.Asset{
			Step:    step,
			Content: asset.Content,
		})
	}

	result := ticketdto.GetTicketByIDResp{
		Message: "",
		Ticket: ticketdto.Ticket{
			ID:          tkt.ID,
			Title:       tkt.Title,
			Description: tkt.Description,
			Assets:      assetsDTO,
		},
	}

	if err == asset.ErrJobTypeNotFound {
		result.Message = "job_type not found in template variables; no assets generated"
	}
	c.JSON(http.StatusOK, result)
}
