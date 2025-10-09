package json

type hostData struct {
	ID      int    `json:"id"`
	Address string `json:"address"`
	Name    string `json:"name"`
	Port    int    `json:"port"`
	User    string `json:"user"`
}
