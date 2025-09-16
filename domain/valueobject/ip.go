package valueobject

import (
	"errors"
	"fmt"
	"net"
)

type IP string

func NewIP(ip string) (IP, error) {
	if net.ParseIP(ip) == nil {
		err := errors.New(fmt.Sprintf("Invalid IP address format: %s", ip))
		return "", err
	}
	return IP(ip), nil
}
