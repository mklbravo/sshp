package ports

import "github.com/mklbravo/sshp/domain/entity"

type ConnectionService interface {
	ConnectToHost(profile *entity.Profile) error
}
