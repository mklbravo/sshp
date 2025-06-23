package main

import (
	"github.com/mklbravo/sshp/application"
)

func main() {
	// SSH connection parameters
	user := "TODO"
	password := "TODO"
	host := "TODO"

	application.ConnectToSSHHost(user, password, host)
}
