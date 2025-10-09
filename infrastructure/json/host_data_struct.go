package json

type hostData struct {
	Address string `json:"address"`
	Group   string `json:"group,omitempty"`
	Name    string `json:"name"`
	Port    int    `json:"port"`
	User    string `json:"user"`
}
