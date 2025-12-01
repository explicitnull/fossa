package ticketrepo

import (
	"context"
	"database/sql"
	"fmt"

	"fossa/service/ticket"

	sq "github.com/Masterminds/squirrel"
	"github.com/pkg/errors"
)

const (
	ticketsTable = "tickets"
)

type SQLite struct {
	database *sql.DB
}

func NewSQLite(database *sql.DB) *SQLite {
	return &SQLite{
		database: database,
	}
}

func (s *SQLite) FetchTickets(ctx context.Context) ([]ticket.Ticket, error) {
	query, values, _ := sq.
		Select(
			"id",
			"title",
			"description",
			"created_at",
		).
		From(ticketsTable).
		OrderByClause(fmt.Sprintf("%s %s", "id", "ASC")).
		PlaceholderFormat(sq.Dollar).
		ToSql()

	results := []ticket.Ticket{}

	rows, err := s.database.QueryContext(ctx, query, values...)
	if err != nil {
		return nil, errors.Wrap(err, "cannot perform select query")
	}

	defer rows.Close()

	for rows.Next() {
		var t ticket.Ticket

		if err := rows.Scan(
			&t.ID,
			&t.Title,
			&t.Description,
			&t.CreatedAt,
		); err != nil {
			return nil, errors.Wrap(err, "cannot perform select query")
		}

		results = append(results, t)
	}

	return results, nil
}
