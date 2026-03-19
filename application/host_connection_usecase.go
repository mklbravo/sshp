package application

import (
	"github.com/mklbravo/sshp/application/ports"
	"github.com/mklbravo/sshp/domain/entity"
)

type HostConnectionUseCase struct {
	conn ports.ConnectionService
}

func NewHostConnectionUseCase(
	conn ports.ConnectionService,
) *HostConnectionUseCase {
	return &HostConnectionUseCase{
		conn: conn,
	}
}

func (this *HostConnectionUseCase) Execute(profile *entity.Profile) error {
	return this.conn.ConnectToHost(profile)
}
