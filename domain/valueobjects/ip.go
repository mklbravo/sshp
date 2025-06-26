package valueobjects

import (
    "errors"
    "net"
)

var ErrInvalidIP = errors.New("invalid IP address")

type IP string

func NewIP(ip string) (IP, error) {
    if net.ParseIP(ip) == nil {
        return "", ErrInvalidIP
    }
    return IP(ip), nil
}
