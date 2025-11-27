package history

import "time"

type HistoryData struct {
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
	CreatedBy string
	UpdatedBy string
	DeletedBy string
}
