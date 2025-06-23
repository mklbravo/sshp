package application

import (
	"github.com/mklbravo/sshp/domain"
	"github.com/mklbravo/sshp/infrastructure"
)

func ConnectToSSHHost(host *domain.Host) {
	sshSession, _ := infrastructure.StartSSHSession(host)
	infrastructure.RunSSHShell(sshSession)
	defer sshSession.Close()
}
