package assetdto

import "fossa/pkg/history"

type Asset struct {
	JobType             string `json:"job_type"`
	Step                string `json:"step"`
	Content             string `json:"content"`
	GenericTemplateUsed bool   `json:"generic_template_used"`
	history.HistoryData
}
