package entity

import (
	"fmt"

	"github.com/mklbravo/sshp/domain/valueobject"
)

type Host struct {
	Group    string
	ID       int
	IP       valueobject.IP
	Name     valueobject.HostName
	Port     valueobject.Port
	Username valueobject.Username
	Details  []string
}

func NewHost(
	id int,
	name string,
	username string,
	ip string,
	port int,
	group string,
	details []string,
) (*Host, error) {
	hostName := valueobject.HostName(name)

	user := valueobject.Username(username)

	hostIP, err := valueobject.NewIP(ip)
	if err != nil {
		return nil, fmt.Errorf("Invalid IP address: %w", err)
	}

	hostPort, err := valueobject.NewPort(port)
	if err != nil {
		return nil, fmt.Errorf("Invalid port: %w", err)
	}

	return &Host{
		Group:    group,
		ID:       id,
		IP:       hostIP,
		Name:     hostName,
		Port:     hostPort,
		Username: user,
		Details:  details,
	}, nil
}

func (this *Host) GetFullAddress() string {
	return fmt.Sprintf("%s:%d",
		string(this.IP),
		int(this.Port),
	)
}

func (this *Host) GetDetailsString() string {
	details := ""
	for _, detail := range this.Details {
		details += detail + " "
	}
	return details
}
