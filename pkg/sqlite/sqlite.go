package sqlite

import (
	"database/sql"
	"fmt"

	_ "github.com/ncruces/go-sqlite3/driver"
	_ "github.com/ncruces/go-sqlite3/embed"
)

func NewDB() (*sql.DB, error) {
	var version string

	db, err := sql.Open("sqlite3", "file:fossa.db")
	if err != nil {
		return nil, err
	}

	db.QueryRow(`SELECT sqlite_version()`).Scan(&version)

	fmt.Printf("db version: %s", version)

	return db, err
}
