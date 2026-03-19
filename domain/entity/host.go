package entity

import (
	"fmt"

	"github.com/mklbravo/sshp/domain/valueobject"
)

type Profile struct {
	Group    string
	ID       int
	IP       valueobject.IP
	Name     valueobject.HostName
	Port     valueobject.Port
	Username valueobject.Username
	Details  []string
}

func NewProfile(
	id int,
	name string,
	username string,
	ip string,
	port int,
	group string,
	details []string,
) (*Profile, error) {
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

	return &Profile{
		Group:    group,
		ID:       id,
		IP:       hostIP,
		Name:     hostName,
		Port:     hostPort,
		Username: user,
		Details:  details,
	}, nil
}

func (this *Profile) GetFullAddress() string {
	return fmt.Sprintf("%s:%d",
		string(this.IP),
		int(this.Port),
	)
}

func (this *Profile) HasDetails() bool {
	return len(this.Details) > 0
}
