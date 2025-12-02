package httpserver

import (
	"fossa/pkg/templatedto"
	"net/http"

	"github.com/gin-gonic/gin"
)

// func (s *Server) GetTemplateByID(c *gin.Context) {
// 	ctx := c.Request.Context()

// 	templateID := c.Param("id")

// 	tpl, err := s.templateService.GetTemplateByID(ctx, templateID)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, c.Error(err))

// 		return
// 	}

// 	result := templatedto.GetTemplateByIDResp{
// 		Message: "",
// 		Template: templatedto.Template{
// 			ID:      tpl.ID,
// 			JobType: tpl.JobType,
// 			Step:    tpl.Step,
// 			Content: tpl.Content,
// 		},
// 	}

// 	if err == asset.ErrJobTypeNotFound {
// 		result.Message = "job_type not found in template variables; no assets generated"
// 	}

// 	c.JSON(http.StatusOK, result)
// }

func (s *Server) GetTemplatesByJobType(c *gin.Context) {
	ctx := c.Request.Context()

	jobType := c.Param("id")

	tpl, err := s.templateService.FetchTemplatesByJobType(ctx, jobType)
	if err != nil {
		c.JSON(http.StatusInternalServerError, c.Error(err))

		return
	}

	if len(tpl) == 0 {
		c.JSON(http.StatusBadRequest, templatedto.GetTemplatesByJobTypeResp{
			Message: "no templates found for the given job_type",
		})
		return
	}

	tplsDTO := make([]templatedto.Template, 0, len(tpl))

	for _, t := range tpl {
		templateDTO := templatedto.Template{
			ID:      t.ID,
			JobType: t.JobType,
			Step:    t.Step,
			Content: t.Content,
		}

		tplsDTO = append(tplsDTO, templateDTO)
	}
	result := templatedto.GetTemplatesByJobTypeResp{
		Message:   "",
		Templates: tplsDTO,
	}

	c.JSON(http.StatusOK, result)
}
