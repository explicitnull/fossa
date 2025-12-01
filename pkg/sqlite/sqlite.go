package sqlite

import (
	"database/sql"

	_ "github.com/ncruces/go-sqlite3/driver"
	_ "github.com/ncruces/go-sqlite3/embed"
)

func NewDB() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "file:db/fossa.db")
	if err != nil {
		return nil, err
	}

	return db, err
}
