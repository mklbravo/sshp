package entity

import (
	"fmt"

	"github.com/mklbravo/sshp/domain/valueobject"
)

type Host struct {
	ID       int
	Name     valueobject.HostName
	Username valueobject.Username
	IP       valueobject.IP
	Port     valueobject.Port
}

func NewHost(id int, name, username, ip string, port int) (*Host, error) {
	hostName := valueobject.HostName(name)

	user := valueobject.Username(username)

	hostIP, err := valueobject.NewIP(ip)
	if err != nil {
		return nil, fmt.Errorf("invalid IP address: %w", err)
	}

	hostPort, err := valueobject.NewPort(port)
	if err != nil {
		return nil, fmt.Errorf("invalid port: %w", err)
	}

	return &Host{
		ID:       id,
		Name:     hostName,
		Username: user,
		IP:       hostIP,
		Port:     hostPort,
	}, nil
}

func (this *Host) GetFullAddress() string {
	return fmt.Sprintf("%s:%d",
		string(this.IP),
		int(this.Port),
	)
}
