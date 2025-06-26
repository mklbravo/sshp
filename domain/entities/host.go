package entities

type Host struct {
    ID   string
    Name valueobjects.HostName
    IP   valueobjects.IP
    Port valueobjects.Port
}
