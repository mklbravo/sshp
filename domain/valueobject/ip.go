package valueobject

import (
	"fmt"
	"net"
)

type IP string

func NewIP(ip string) (IP, error) {
	if net.ParseIP(ip) == nil {
		err := fmt.Errorf("Invalid IP address format: %s", ip)
		return "", err
	}
	return IP(ip), nil
}
