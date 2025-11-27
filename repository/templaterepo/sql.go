package templaterepo

import (
	"context"
	"database/sql"
	"fmt"

	"fossa/service/template"

	sq "github.com/Masterminds/squirrel"
	"github.com/pkg/errors"
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

func (s *SQLite) FetchTemplatesByName(ctx context.Context, name string) ([]template.Template, error) {
	query, values, _ := sq.
		Select(
			"id",
			"name",
			"step",
			"contents",
		).
		From(templatesTable).
		OrderByClause(fmt.Sprintf("%s %s", "name", "ASC")).
		PlaceholderFormat(sq.Dollar).
		ToSql()

	results := []template.Template{}

	rows, err := s.database.QueryContext(ctx, query, values...)
	if err != nil {
		return nil, errors.Wrap(err, "cannot perform select query")
	}

	defer rows.Close()

	for rows.Next() {
		var t template.Template

		if err := rows.Scan(
			&t.ID,
			&t.Name,
			&t.Step,
			&t.Content,
		); err != nil {
			return nil, errors.Wrap(err, "cannot perform select query")
		}

		results = append(results, t)
	}

	return results, nil
}
