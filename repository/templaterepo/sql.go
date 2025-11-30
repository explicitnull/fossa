package templaterepo

import (
	"context"
	"database/sql"

	"fossa/service/template"
)

const (
	templatesTable = "templates"
)

type SQLite struct {
	database *sql.DB
}

func NewSQLite(database *sql.DB) *SQLite {
	return &SQLite{
		database: database,
	}
}

func (s *SQLite) FetchTemplatesByJobType(ctx context.Context, jobType string) ([]template.Template, error) {
	// query, values, _ := sq.
	// 	Select(
	// 		"id",
	// 		"job_type",
	// 		"step",
	// 		"contents",
	// 	).
	// 	From(templatesTable).
	// 	OrderByClause(fmt.Sprintf("%s %s", "job_type", "ASC")).
	// 	PlaceholderFormat(sq.Dollar).
	// 	ToSql()

	results := []template.Template{}

	// rows, err := s.database.QueryContext(ctx, query, values...)
	// if err != nil {
	// 	return nil, errors.Wrap(err, "cannot perform select query")
	// }

	// defer rows.Close()

	// for rows.Next() {
	// 	var t template.Template

	// 	if err := rows.Scan(
	// 		&t.ID,
	// 		&t.JobType,
	// 		&t.Step,
	// 		&t.Content,
	// 	); err != nil {
	// 		return nil, errors.Wrap(err, "cannot perform select query")
	// 	}

	// 	results = append(results, t)
	// }

	// TODO: remove mock data

	t := template.Template{
		ID:      "1",
		JobType: "install_optics_and_connect_two_cenic_devices",
		Step:    "installation",
		Content: "example-content:\n A device: {{ device_a_port }}",
	}

	results = append(results, t)

	return results, nil
}
