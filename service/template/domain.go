package template

import (
	"fossa/pkg/history"
)

const (
	AddressConfirmation  = "10 address_confirmation"
	Shipping             = "20 shipping"
	InboundShipping      = "30 inbound_shipping"
	Installation         = "40 installation"
	InstallationMOP      = "42 installation_mop"
	MaintenanceCalendar  = "50 maintenance_calendar"
	Migration            = "60 migration"
	Documentation        = "70 documentation"
	Readiness            = "80 readiness"
	SoftwareUploadAndTAC = "87 software_upload_and_jtac"
	Decommission         = "100 decommission"
)

var Steps = []string{
	AddressConfirmation,
	Shipping,
	InboundShipping,
	Installation,
	InstallationMOP,
	MaintenanceCalendar,
	Migration,
	Documentation,
	Readiness,
	SoftwareUploadAndTAC,
	Decommission,
}

var HumanReadableSteps = map[string]string{
	AddressConfirmation:  "Address confirmation e-mail",
	Shipping:             "Shipping ticket",
	InboundShipping:      "Inbound shipping ticket or email",
	Installation:         "Installation ticket",
	InstallationMOP:      "Installation MOP",
	MaintenanceCalendar:  "Maintenance Calendar event",
	Migration:            "Migration MOP",
	Documentation:        "Documentation in Cobber, Netbox and Confluence",
	Readiness:            "Operational Readiness Ticket",
	SoftwareUploadAndTAC: "Ticket for uploading new software to device and opening a TAC case",
	Decommission:         "Decommissioning ticket",
}

type Template struct {
	ID                  string
	JobType             string
	Step                string
	Content             string
	GenericTemplateUsed bool
	history.HistoryData
}
