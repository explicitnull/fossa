package httpserver

import (
	"errors"
	"fossa/pkg/dto/assetdto"
	"fossa/pkg/ticketdto"
	"fossa/service/asset"
	"fossa/service/template"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *Server) GetTickets(c *gin.Context) {
	ctx := c.Request.Context()

	assignedTo, err := mapAssignee(c.Query("assigned_to"))
	if err != nil {
		c.JSON(http.StatusBadRequest, c.Error(err))

		return
	}

	tickets, err := s.ticketService.GetTickets(ctx, assignedTo)
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

	ticketID := c.Query("key")

	var varsFromFile bool
	if c.Query("from_file") == "yes" {
		varsFromFile = true
	}

	tkt, err := s.ticketService.GetTicketByID(ctx, ticketID, varsFromFile)
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
	assetsDTO := make([]assetdto.Asset, 0, len(assets))

	for _, stepName := range template.Steps {
		stepAsset, exist := assets[stepName]
		if !exist {
			continue
		}

		assetsDTO = append(assetsDTO, assetdto.Asset{
			Step:    template.HumanReadableSteps[stepName],
			Content: stepAsset.Content,
		})
	}

	result := ticketdto.GetTicketByIDResp{
		Message: "",
		Ticket: ticketdto.Ticket{
			ID:          tkt.ID,
			Title:       tkt.Title,
			Description: tkt.Description,
		},
		Assets: assetsDTO,
	}

	if err == asset.ErrJobTypeNotFound {
		result.Message = "job_type not found in template variables; no assets generated"
	}
	c.JSON(http.StatusOK, result)
}

func mapAssignee(URLparam string) (string, error) {
	switch URLparam {
	case "Abdel":
		return "Abdelrahman Mohammad", nil
	case "Alaa":
		return "Alaa Shaheen", nil
	case "Dmitrii":
		return "Dmitrii Zhargalov", nil
	case "Jose":
		return "Jose Gonzalez", nil
	case "":
		return "", nil
	}

	return "", errors.New("assignee name not found")
}
