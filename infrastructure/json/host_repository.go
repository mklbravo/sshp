package json

import (
	"encoding/json"
	"os"

	"github.com/mklbravo/sshp/domain/entity"
)

type JsonHostRepository struct {
	allHosts     []*entity.Host
	indexedHosts map[int]*entity.Host
}

func NewJsonHostRepository(filePath string) (*JsonHostRepository, error) {
	fileHosts, _ := loadHostsFromFile(filePath)

	var allHosts []*entity.Host
	var indexedHosts = make(map[int]*entity.Host)

	for _, hd := range fileHosts {
		host, err := entity.NewHost(hd.ID, hd.Name, hd.User, hd.Address, hd.Port)
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

func loadHostsFromFile(filePath string) ([]*hostData, error) {
	fileData, err := os.ReadFile(filePath)
	if err != nil {
		file, err := os.Create(filePath)
		if err != nil {
			return nil, err
		}
		file.Write([]byte("[]"))
		defer file.Close()

		return []*hostData{}, nil
	}

	var jsonHosts []*hostData
	err = json.Unmarshal(fileData, &jsonHosts)
	if err != nil {
		return nil, err
	}

	return jsonHosts, nil
}

func (this *JsonHostRepository) FindByID(id int) (*entity.Host, error) {
	return this.indexedHosts[id], nil
}
func (this *JsonHostRepository) FindAll() ([]*entity.Host, error) {
	return this.allHosts, nil
}

func (this *JsonHostRepository) Save(host *entity.Host) error {
	// TODO
	return nil
}
