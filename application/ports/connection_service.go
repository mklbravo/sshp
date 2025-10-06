package ports

import "github.com/mklbravo/sshp/domain/entity"

type ConnectionService interface {
	ConnectToHost(host *entity.Host) error
}
