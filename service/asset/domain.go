package asset

import "fossa/pkg/history"

var Steps = map[string]struct{}{
	"10 address_confirmation":     {},
	"20 shipping":                 {},
	"30 inbound_shipping":         {},
	"40 installation":             {},
	"42 installation_mop":         {},
	"50 conf_calendar":            {},
	"60 migration":                {},
	"70 documentation":            {},
	"80 readiness":                {},
	"87 software_upload_and_jtac": {},
	"100 decommission":            {},
}

// business domain entities
type Asset struct {
	ID      string
	JobType string
	Step    string
	Content string
	history.HistoryData
}
