package application

import (
	"github.com/mklbravo/sshp/domain/entity"
	"github.com/mklbravo/sshp/domain/repository"
)

type HostListUseCase struct {
	HostRepository repository.HostRepository
}

func NewHostListUseCase(hostRepo repository.HostRepository) *HostListUseCase {
	return &HostListUseCase{
		HostRepository: hostRepo,
	}
}

func (uc *HostListUseCase) Execute() ([]*entity.Profile, error) {
	hostList, err := uc.HostRepository.FindAll()
	if err != nil {
		return nil, err
	}

	return hostList, nil
}
