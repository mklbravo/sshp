package valueobjects

import "errors"

var ErrInvalidPort = errors.New("invalid port")

type Port int

func NewPort(port int) (Port, error) {
    if port < 1 || port > 65535 {
        return 0, ErrInvalidPort
    }
    return Port(port), nil
}
