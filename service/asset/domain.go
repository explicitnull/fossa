package asset

import "fossa/pkg/history"

// business domain entities
type Asset struct {
	ID      string
	JobType string
	Step    string
	Content string
	history.HistoryData
}
