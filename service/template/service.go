package template

import (
	"context"
	"fossa/pkg/history"

	"github.com/pkg/errors"
)

// business domain entities
type Template struct {
	ID      string
	Name    string
	Step    string
	Content string
	history.HistoryData
}

// end of entities

type Service struct {
	repository TemplateRepository
}

func NewService(
	repository TemplateRepository,
) *Service {
	return &Service{
		repository: repository,
	}
}

func (s *Service) FetchTemplates(ctx context.Context, name string) ([]Template, error) {
	templates, err := s.repository.FetchTemplatesByName(ctx, name)
	if err != nil {
		return nil, errors.Wrap(err, "can't get templates")
	}

	return templates, nil
}
