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

// scanHostRow scans a single row (from QueryRow) into a Host entity.
func scanHostRow(scanner interface {
    Scan(dest ...interface{}) error
}) (*entities.Host, error) {
    var id, name, ip string
    var port int
    if err := scanner.Scan(&id, &name, &ip, &port); err != nil {
        return nil, err
    }
    ipVO, err := valueobjects.NewIP(ip)
    if err != nil {
        return nil, err
    }
    portVO, err := valueobjects.NewPort(port)
    if err != nil {
        return nil, err
    }
    return &entities.Host{
        ID:   id,
        Name: valueobjects.HostName(name),
        IP:   ipVO,
        Port: portVO,
    }, nil
}


func NewSqliteHostRepository(db *sql.DB) repository.HostRepository {
    return &SqliteHostRepository{db: db}
}

func (r *SqliteHostRepository) FindByID(id string) (*entities.Host, error) {
    row := r.db.QueryRow("SELECT id, name, ip, port FROM hosts WHERE id = ?", id)
    return scanHostRow(row)
}

func (r *SqliteHostRepository) Save(host *entities.Host) error {
    _, err := r.db.Exec(
        "INSERT OR REPLACE INTO hosts (id, name, ip, port) VALUES (?, ?, ?, ?)",
        host.ID, string(host.Name), string(host.IP), int(host.Port),
    )
    return err
}

func (r *SqliteHostRepository) FindAll() ([]*entities.Host, error) {
    rows, err := r.db.Query("SELECT id, name, ip, port FROM hosts")
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var hosts []*entities.Host
    for rows.Next() {
        host, err := scanHostRow(rows)
        if err != nil {
            return nil, err
        }
        hosts = append(hosts, host)
    }
    if err := rows.Err(); err != nil {
        return nil, err
    }
    return hosts, nil
}
