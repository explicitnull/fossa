package template

import (
	"context"
)

type TemplateRepository interface {
	FetchTemplatesByName(ctx context.Context, name string) ([]Template, error)
}
