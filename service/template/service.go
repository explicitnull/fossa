package template

import (
	"context"

	"github.com/pkg/errors"
)

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

func (s *Service) FetchTemplatesByJobType(ctx context.Context, jobType string) ([]Template, error) {
	templates, err := s.repository.FetchTemplatesByJobType(ctx, jobType)
	if err != nil {
		return nil, errors.Wrap(err, "can't get templates")
	}

	return templates, nil
}
