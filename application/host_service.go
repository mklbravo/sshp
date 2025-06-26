package application

import (
	"github.com/mklbravo/sshp/domain"
	"github.com/mklbravo/sshp/infrastructure"
)

func GetAllHosts() []domain.Host {
	return infrastructure.GetHostFromMemory()
}
