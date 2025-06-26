package infrastructure

import (
    "database/sql"
    "github.com/mklbravo/sshp/domain/entities"
    "github.com/mklbravo/sshp/domain/repository"
    "github.com/mklbravo/sshp/domain/valueobjects"
    _ "github.com/mattn/go-sqlite3"
)

type SqliteHostRepository struct {
    db *sql.DB
}

func NewSqliteHostRepository(db *sql.DB) repository.HostRepository {
    return &SqliteHostRepository{db: db}
}

func (r *SqliteHostRepository) FindByID(id string) (*entities.Host, error) {
    row := r.db.QueryRow("SELECT id, name FROM hosts WHERE id = ?", id)
    var host entities.Host
    var name string
    if err := row.Scan(&host.ID, &name); err != nil {
        return nil, err
    }
    host.Name = valueobjects.HostName(name)
    return &host, nil
}

func (r *SqliteHostRepository) Save(host *entities.Host) error {
    _, err := r.db.Exec(
        "INSERT OR REPLACE INTO hosts (id, name) VALUES (?, ?)",
        host.ID, string(host.Name),
    )
    return err
}
