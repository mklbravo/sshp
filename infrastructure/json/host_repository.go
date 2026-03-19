package json

import (
	"encoding/json"
	"os"

	"github.com/mklbravo/sshp/domain/entity"
)

type JsonHostRepository struct {
	allHosts     []*entity.Profile
	indexedHosts map[int]*entity.Profile
}

func NewJsonHostRepository(filePath string) (*JsonHostRepository, error) {
	fileHosts, _ := loadHostsFromFile(filePath)

	var allHosts []*entity.Profile
	var indexedHosts = make(map[int]*entity.Profile)

	for index, hd := range fileHosts {
		if hd.Group == "" {
			hd.Group = "default"
		}

		host, err := entity.NewProfile(
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

		allHosts = append(allHosts, host)
		indexedHosts[host.ID] = host
	}

	return &JsonHostRepository{
		allHosts:     allHosts,
		indexedHosts: indexedHosts,
	}, nil
}

func loadHostsFromFile(filePath string) ([]*profileDTO, error) {
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

func (this *JsonHostRepository) FindByID(id int) (*entity.Profile, error) {
	return this.indexedHosts[id], nil
}
func (this *JsonHostRepository) FindAll() ([]*entity.Profile, error) {
	return this.allHosts, nil
}

func (this *JsonHostRepository) Save(profile *entity.Profile) error {
	// TODO
	return nil
}
