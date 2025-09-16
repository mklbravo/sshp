package sqlite

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
	"github.com/mklbravo/sshp/domain/entity"
)

// This defines an interface for use in the scanHostRow function
// This allows us to use both sql.Row and sql.Rows as they both implement the Scan method
type rowScanner interface {
	Scan(dest ...any) error
}

type SqliteHostRepository struct {
	db *sql.DB
}

func scanHostRow(scanner rowScanner) (*entity.Host, error) {
	var id, port int
	var name, username, ip string

	err := scanner.Scan(&id, &name, &username, &ip, &port)
	if err != nil {
		return nil, err
	}

	host, err := entity.NewHost(id, name, username, ip, port)

	if err != nil {
		return nil, err
	}

	return host, nil
}

func NewHostRepository(db *sql.DB) *SqliteHostRepository {
	return &SqliteHostRepository{db: db}
}

func (r *SqliteHostRepository) FindByID(id string) (*entity.Host, error) {
	row := r.db.QueryRow("SELECT id, name, username, ip, port FROM hosts WHERE id = ?", id)
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
	rows, err := r.db.Query("SELECT id, name, username, ip, port FROM hosts")
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
