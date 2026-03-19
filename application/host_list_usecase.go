package application

import (
	"github.com/mklbravo/sshp/domain/entity"
	"github.com/mklbravo/sshp/domain/repository"
)

type HostListUseCase struct {
	ProfileRepository repository.ProfileRepository
}

func NewHostListUseCase(profileRepository repository.ProfileRepository) *HostListUseCase {
	return &HostListUseCase{
		ProfileRepository: profileRepository,
	}
}

func (uc *HostListUseCase) Execute() ([]*entity.Profile, error) {
	profileList, err := uc.ProfileRepository.FindAll()
	if err != nil {
		return nil, err
	}

	return profileList, nil
}
