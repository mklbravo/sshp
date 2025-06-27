package entity

import (
	"fmt"

	"github.com/mklbravo/sshp/domain/valueobject"
)

type Host struct {
	ID       string
	Name     valueobject.HostName
	Username valueobject.Username
	IP       valueobject.IP
	Port     valueobject.Port
}

func (this *Host) GetFullAddress() string {
	return fmt.Sprintf("%s:%d",
		string(this.IP),
		int(this.Port),
	)
}
