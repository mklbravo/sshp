package infrastructure

import (
    "database/sql"
    "github.com/mklbravo/sshp/domain/repository"
    _ "github.com/mattn/go-sqlite3"
	"github.com/mklbravo/sshp/domain/entity"
	"github.com/mklbravo/sshp/domain/valueobject"
)

type SqliteHostRepository struct {
    db *sql.DB
}

// scanHostRow scans a single row (from QueryRow) into a Host entity.
func scanHostRow(scanner interface {
    Scan(dest ...interface{}) error
    var id, name, ip string
    var port int
    if err := scanner.Scan(&id, &name, &ip, &port); err != nil {
        return nil, err
    }
    if err != nil {
        return nil, err
    }
    if err != nil {
        return nil, err
    }
        ID:   id,
        IP:   ipVO,
        Port: portVO,
    }, nil
}) (*entity.Host, error) {
	ipVO, err := valueobject.NewIP(ip)
	portVO, err := valueobject.NewPort(port)
	return &entity.Host{
		Name: valueobject.HostName(name),
}


func NewSqliteHostRepository(db *sql.DB) repository.HostRepository {
    return &SqliteHostRepository{db: db}
}

    row := r.db.QueryRow("SELECT id, name, ip, port FROM hosts WHERE id = ?", id)
    return scanHostRow(row)
func (r *SqliteHostRepository) FindByID(id string) (*entity.Host, error) {
}

    _, err := r.db.Exec(
        "INSERT OR REPLACE INTO hosts (id, name, ip, port) VALUES (?, ?, ?, ?)",
        host.ID, string(host.Name), string(host.IP), int(host.Port),
    )
    return err
func (r *SqliteHostRepository) Save(host *entity.Host) error {
}

    rows, err := r.db.Query("SELECT id, name, ip, port FROM hosts")
    if err != nil {
        return nil, err
    }
    defer rows.Close()
func (r *SqliteHostRepository) FindAll() ([]*entity.Host, error) {

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
	var hosts []*entity.Host
}
