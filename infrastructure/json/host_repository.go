package json

import (
	"encoding/json"
	"os"

	"github.com/mklbravo/sshp/domain/entity"
)

type JsonProfileRepository struct {
	allProfiles     []*entity.Profile
	indexedProfiles map[int]*entity.Profile
}

func NewJsonProfileRepository(filePath string) (*JsonProfileRepository, error) {
	fileProfiles, _ := loadProfilesFromFile(filePath)

	var allProfiles []*entity.Profile
	var indexedProfiles = make(map[int]*entity.Profile)

	for index, hd := range fileProfiles {
		if hd.Group == "" {
			hd.Group = "default"
		}

		profile, err := entity.NewProfile(
			index,
			hd.Name,
			hd.User,
			hd.Address,
			hd.Port,
			hd.Group,
			hd.Details,
		)

		if err != nil {
			return nil, err
		}

		allProfiles = append(allProfiles, profile)
		indexedProfiles[profile.ID] = profile
	}

	return &JsonProfileRepository{
		allProfiles:     allProfiles,
		indexedProfiles: indexedProfiles,
	}, nil
}

func loadProfilesFromFile(filePath string) ([]*profileDTO, error) {
	fileData, err := os.ReadFile(filePath)
	if err != nil {
		file, err := os.Create(filePath)
		if err != nil {
			return nil, err
		}
		file.Write([]byte("[]"))
		defer file.Close()

		return []*profileDTO{}, nil
	}

	var jsonProfiles []*profileDTO
	err = json.Unmarshal(fileData, &jsonProfiles)
	if err != nil {
		return nil, err
	}

	return jsonProfiles, nil
}

func (this *JsonProfileRepository) FindByID(id int) (*entity.Profile, error) {
	return this.indexedProfiles[id], nil
}
func (this *JsonProfileRepository) FindAll() ([]*entity.Profile, error) {
	return this.allProfiles, nil
}

func (this *JsonProfileRepository) Save(profile *entity.Profile) error {
	// TODO
	return nil
}
