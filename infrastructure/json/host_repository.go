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

func loadHostsFromFile(filePath string) ([]*JsonHostData, error) {
	fileData, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	var jsonHosts []*JsonHostData
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

type JsonHostData struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	User    string `json:"user"`
	Address string `json:"address"`
	Port    int    `json:"port"`
}
