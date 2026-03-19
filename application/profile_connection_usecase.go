package application

import (
	"github.com/mklbravo/sshp/application/ports"
	"github.com/mklbravo/sshp/domain/entity"
)

type ProfileConnectionUseCase struct {
	conn ports.ConnectionService
}

func NewProfileConnectionUseCase(
	conn ports.ConnectionService,
) *ProfileConnectionUseCase {
	return &ProfileConnectionUseCase{
		conn: conn,
	}
}

func (this *ProfileConnectionUseCase) Execute(profile *entity.Profile) error {
	return this.conn.ConnectToHost(profile)
}
