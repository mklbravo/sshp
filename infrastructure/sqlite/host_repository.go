package sqlite

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
	"github.com/mklbravo/sshp/domain/entity"
	"github.com/mklbravo/sshp/domain/repository"
	"github.com/mklbravo/sshp/domain/valueobject"
)

type SqliteHostRepository struct {
	db *sql.DB
}

// scanHostRow scans a single row (from QueryRow) into a Host entity.
func scanHostRow(scanner interface {
	Scan(dest ...interface{}) error
}) (*entity.Host, error) {
	var name, ip string
	var id int
	var port int
	if err := scanner.Scan(&id, &name, &ip, &port); err != nil {
		return nil, err
	}
	ipVO, err := valueobject.NewIP(ip)
	if err != nil {
		return nil, err
	}
	portVO, err := valueobject.NewPort(port)
	if err != nil {
		return nil, err
	}
	return &entity.Host{
		ID:   id,
		Name: valueobject.HostName(name),
		IP:   ipVO,
		Port: portVO,
	}, nil
}

func NewHostRepository(db *sql.DB) *SqliteHostRepository {
	return &SqliteHostRepository{db: db}
}

func (r *SqliteHostRepository) FindByID(id string) (*entity.Host, error) {
	row := r.db.QueryRow("SELECT id, name, ip, port FROM hosts WHERE id = ?", id)
	return scanHostRow(row)
}

func (r *SqliteHostRepository) Save(host *entity.Host) error {
	_, err := r.db.Exec(
		"INSERT OR REPLACE INTO hosts (id, name, ip, port) VALUES (?, ?, ?, ?)",
		host.ID, string(host.Name), string(host.IP), int(host.Port),
	)
	return err
}

func (r *SqliteHostRepository) FindAll() ([]*entity.Host, error) {
	rows, err := r.db.Query("SELECT id, name, ip, port FROM hosts")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var hosts []*entity.Host
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
