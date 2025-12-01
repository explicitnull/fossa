package asset

import (
	"context"
	"errors"
	"fmt"
	"fossa/service/template"

	"github.com/nikolalohinski/gonja/v2"
	"github.com/nikolalohinski/gonja/v2/exec"
)

var ErrJobTypeNotFound = errors.New("job_type not found in template variables")

type Service struct {
	repository      AssetRepository
	templateService *template.Service
}

func NewService(repository AssetRepository, templateService *template.Service) *Service {
	return &Service{
		repository:      repository,
		templateService: templateService}
}

func (s *Service) GenerateAssetsForTicket(ctx context.Context, vars map[string]interface{}) (map[string]Asset, error) {
	// fetch templates

	jobTypeI, ok := vars["job_type"]
	if !ok {
		return nil, ErrJobTypeNotFound
	}

	jobType, ok := jobTypeI.(string)
	if !ok {
		return nil, fmt.Errorf("job_type is not a string")
	}

	templates, err := s.templateService.FetchTemplatesByJobType(ctx, jobType)
	if err != nil {
		return nil, err
	}

	// generate assets

	res := make(map[string]Asset)

	for _, tpl := range templates {
		generated, err := s.executeTemplate(ctx, tpl.Content, vars)
		if err != nil {
			return nil, err
		}

		res[tpl.Step] = Asset{
			Content: generated,
		}
	}

	return res, nil
}

func (s *Service) executeTemplate(ctx context.Context, tplContent string, vars map[string]interface{}) (string, error) {
	tpl, err := gonja.FromString(tplContent)
	if err != nil {
		return "", err
	}

	data := exec.NewContext(vars)

	return tpl.ExecuteToString(data)
}
