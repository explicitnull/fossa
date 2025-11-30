package template

import (
	"context"
)

type TemplateRepository interface {
	FetchTemplatesByJobType(ctx context.Context, jobType string) ([]Template, error)
}
