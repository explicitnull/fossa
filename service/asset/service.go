package asset

import (
	"context"
	"fossa/service/template"
)

type ticketAssets map[string]Asset // map key: step

type Service struct {
	repository      AssetRepository
	templateService *template.Service
}

func NewService(repository AssetRepository, templateService *template.Service) *Service {
	return &Service{
		repository:      repository,
		templateService: templateService}
}

func (s *Service) GenerateAssetsForTicket(ctx context.Context, vars map[string]string) (map[string]Asset, error) {
	res := make(map[string]Asset)

	// for _, step := range vars {

	// }

	return res, nil
}
