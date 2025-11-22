// package ticketrepo

// import (
// 	"context"
// 	"fmt"

// 	"fossa/service/ticket"

// 	"github.com/pkg/errors"

// 	sq "github.com/Masterminds/squirrel"
// )

// const (
// 	settingsTable = "webhook_settings"
// )

// type Postgres struct {
// 	database *postgres.Database
// }

// func NewPostgres(database *postgres.Database) *Postgres {
// 	return &Postgres{
// 		database: database,
// 	}
// }

// func (p *Postgres) ReadSettings(ctx context.Context) ([]ticket.Settings, error) {
// 	query, values, _ := sq.
// 		Select(
// 			"id",
// 			"name",
// 			"channel_type",
// 			"settings",
// 			"created_at",
// 		).
// 		From(settingsTable).
// 		OrderByClause(fmt.Sprintf("%s %s", "name", "ASC")).
// 		PlaceholderFormat(sq.Dollar).
// 		ToSql()

// 	results := []ticket.Settings{}

// 	rows, err := p.database.Query(ctx, query, values...)
// 	if err != nil {
// 		return nil, errors.Wrap(err, "cannot perform select query")
// 	}

// 	defer rows.Close()

// 	for rows.Next() {
// 		var ticket ticket.Settings

// 		if err := rows.Scan(
// 			&ticket.ID,
// 			&ticket.Name,
// 			&ticket.ChannelType,
// 			&ticket.Settings,
// 			&ticket.CreatedAt,
// 		); err != nil {
// 			return nil, errors.Wrap(err, "cannot perform select query")
// 		}

// 		results = append(results, ticket)
// 	}

// 	return results, nil
// }
