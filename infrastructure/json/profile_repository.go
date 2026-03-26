package json

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/mklbravo/sshp/domain/entity"
)

type JsonProfileRepository struct {
	allProfiles     []*entity.Profile
	filePath        string
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
		filePath:        filePath,
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

func (this *JsonProfileRepository) Save(newProfile *entity.Profile) error {
	// Check if profile already exists
	for _, existing := range this.allProfiles {
		if newProfile.IsSame(existing) {
			return fmt.Errorf(
				"Profile with for User '%s' on IP '%s' already exists [%s]",
				existing.Username,
				existing.IP,
				existing.Name,
			)
		}
	}

	newProfile.ID = len(this.allProfiles) // Assign new ID based on current count

	// Add to in-memory collections
	this.allProfiles = append(this.allProfiles, newProfile)
	this.indexedProfiles[newProfile.ID] = newProfile

	// Convert all profiles to DTOs
	var jsonProfiles []*profileDTO
	for _, profile := range this.allProfiles {
		jsonProfiles = append(jsonProfiles, &profileDTO{
			Name:    string(profile.Name),
			User:    string(profile.Username),
			Address: string(profile.IP),
			Port:    int(profile.Port),
			Group:   profile.Group,
			Details: profile.Details,
		})
	}

	// Marshal and save
	data, err := json.MarshalIndent(jsonProfiles, "", "  ")
	if err != nil {
		return err
	}
	if err := os.WriteFile(this.filePath, data, 0664); err != nil {
		return err
	}
	return nil
}
