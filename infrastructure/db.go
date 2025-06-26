package infrastructure

import (
    "database/sql"
    _ "github.com/mattn/go-sqlite3"
)

func InitDB(dataSourceName string) (*sql.DB, error) {
    db, err := sql.Open("sqlite3", dataSourceName)
    if err != nil {
        return nil, err
    }
    _, err = db.Exec(`CREATE TABLE IF NOT EXISTS hosts (
        id TEXT PRIMARY KEY,
        name TEXT NOT NULL,
        ip TEXT NOT NULL,
        port INTEGER NOT NULL
    )`)
    if err != nil {
        return nil, err
    }
    return db, nil
}
