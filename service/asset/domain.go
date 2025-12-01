package asset

import "fossa/pkg/history"

var Steps = map[string]struct{}{
	"shipping":      {},
	"receiving":     {},
	"installation":  {},
	"documentation": {},
}

// business domain entities
type Asset struct {
	ID      string
	JobType string
	Step    string
	Content string
	history.HistoryData
}
