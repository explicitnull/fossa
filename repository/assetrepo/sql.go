package assetrepo

import (
	"context"
	"database/sql"
	"fmt"

	"fossa/service/asset"

	sq "github.com/Masterminds/squirrel"
	"github.com/pkg/errors"
)

const (
	assetsTable = "assets"
)

type SQLite struct {
	database *sql.DB
}

func NewSQLite(database *sql.DB) *SQLite {
	return &SQLite{
		database: database,
	}
}

func (s *SQLite) FetchAssetsByTicketID(ctx context.Context, ticketID string) ([]asset.Asset, error) {
	query, values, _ := sq.
		Select(
			"id",
			"job_type",
			"step",
			"contents",
		).
		From(assetsTable).
		OrderByClause(fmt.Sprintf("%s %s", "job_type", "ASC")).
		PlaceholderFormat(sq.Dollar).
		ToSql()

	results := []asset.Asset{}

	rows, err := s.database.QueryContext(ctx, query, values...)
	if err != nil {
		return nil, errors.Wrap(err, "send query")
	}

	defer rows.Close()

	for rows.Next() {
		var t asset.Asset

		if err := rows.Scan(
			&t.ID,
			&t.JobType,
			&t.Step,
			&t.Content,
		); err != nil {
			return nil, errors.Wrap(err, "iterate rows")
		}

		results = append(results, t)
	}

	return results, nil
}
