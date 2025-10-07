package repository

import "github.com/mklbravo/sshp/domain/entity"

type HostRepository interface {
	FindByID(id int) (*entity.Host, error)
	FindAll() ([]*entity.Host, error)
	Save(host *entity.Host) error
}
