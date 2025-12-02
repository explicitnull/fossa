package templatedto

import "fossa/pkg/history"

type GetTemplatesByJobTypeResp struct {
	Message   string     `json:"message"`
	Templates []Template `json:"templates"`
}

type GetTemplateByIDResp struct {
	Message  string   `json:"message"`
	Template Template `json:"template"`
}

type Template struct {
	ID                  string `json:"id"`
	JobType             string `json:"job_type"`
	Step                string `json:"step"`
	Content             string `json:"content"`
	GenericTemplateUsed bool   `json:"generic_template_used"`
	history.HistoryData
}
