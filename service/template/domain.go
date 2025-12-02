package template

import (
	"fossa/pkg/history"
)

// business domain entities
type Template struct {
	ID                  string
	JobType             string
	Step                string
	Content             string
	GenericTemplateUsed bool
	history.HistoryData
}
