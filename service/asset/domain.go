package asset

import "fossa/pkg/history"

var Steps = map[string]struct{}{
	"01 address_confirmation": {},
	"02 shipping":             {},
	"03 inbound_shipping":     {},
	"04 installation":         {},
	"05 documentation":        {},
	"06 readiness":            {},
	"07 decomission":          {},
}

// business domain entities
type Asset struct {
	ID      string
	JobType string
	Step    string
	Content string
	history.HistoryData
}
