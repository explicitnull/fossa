package ticket

import (
	"fossa/pkg/history"
)

// business domain entities
type Ticket struct {
	ID          string
	Title       string
	Priority    string
	Description string
	Assignee    string
	history.HistoryData

	TemplateVariables map[string]interface{}
}
