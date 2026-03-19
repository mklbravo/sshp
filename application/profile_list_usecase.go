package application

import (
	"github.com/mklbravo/sshp/domain/entity"
	"github.com/mklbravo/sshp/domain/repository"
)

type ProfileListUseCase struct {
	ProfileRepository repository.ProfileRepository
}

func NewProfileListUseCase(profileRepository repository.ProfileRepository) *ProfileListUseCase {
	return &ProfileListUseCase{
		ProfileRepository: profileRepository,
	}
}

func (uc *ProfileListUseCase) Execute() ([]*entity.Profile, error) {
	profileList, err := uc.ProfileRepository.FindAll()
	if err != nil {
		return nil, err
	}

	return profileList, nil
}
