package infrastructure

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

type DatabaseService struct {
	dbConnection   *sql.DB
	dataSourceName string
}

func NewDatabaseService(dataSource string) *DatabaseService {
	dbService := &DatabaseService{
		dataSourceName: dataSource,
	}

	err := dbService.initDB()

	if err != nil {
		panic("Failed to initialize database: " + err.Error())
	}

	return dbService
}

func (this *DatabaseService) initDB() error {
	db, err := sql.Open("sqlite3", this.dataSourceName)
	if err != nil {
		return err
	}
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS hosts (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        ip TEXT NOT NULL,
        name TEXT NOT NULL,
        port INTEGER NOT NULL,
        username TEXT NOT NULL
    )`)

	if err != nil {
		return err
	}
	this.dbConnection = db

	return nil
}
