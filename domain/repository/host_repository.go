package repository

import "github.com/mklbravo/sshp/domain/entity"

type HostRepository interface {
	FindByID(id int) (*entity.Profile, error)
	FindAll() ([]*entity.Profile, error)
	Save(profile *entity.Profile) error
}
