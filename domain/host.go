package domain

import "fmt"

type Host struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Username string `json:"username"`
	IP       string `json:"ip"`
	Port     int    `json:"port"`
}

func (this *Host) GetFullAddress() string {
	return fmt.Sprintf("%s:%d", this.IP, this.Port)
}
