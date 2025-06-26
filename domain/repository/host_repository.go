package repository

import "github.com/mklbravo/sshp/domain/entities"

type HostRepository interface {
    FindByID(id string) (*entities.Host, error)
    Save(host *entities.Host) error
}
