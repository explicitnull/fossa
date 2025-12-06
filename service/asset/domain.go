package asset

import "fossa/pkg/history"

var Steps = map[string]struct{}{
	"01 address_confirmation": {},
	"02 shipping":             {},
	"03 inbound_shipping":     {},
	"04 installation":         {},
	"05 conf_calendar":        {},
	"06 migration":            {},
	"07 documentation":        {},
	"08 readiness":            {},
	"09 decomission":          {},
}

// business domain entities
type Asset struct {
	ID      string
	JobType string
	Step    string
	Content string
	history.HistoryData
}
