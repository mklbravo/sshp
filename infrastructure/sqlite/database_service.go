package sqlite

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

func GetDBConnection(dataSource string) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", dataSource)
	if err != nil {
		return nil, err
	}

	err = migrate(db)

	if err != nil {
		return nil, err
	}

	return db, nil
}

func migrate(db *sql.DB) error {

	_, err := db.Exec(`CREATE TABLE IF NOT EXISTS hosts (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        ip TEXT NOT NULL,
        name TEXT NOT NULL,
        port INTEGER NOT NULL,
        username TEXT NOT NULL
    )`)

	if err != nil {
		return err
	}

	return nil
}
