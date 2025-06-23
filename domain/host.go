package domain

type Host struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Username string `json:"username"`
	IP       string `json:"ip"`
	Port     int    `json:"port"`
}
