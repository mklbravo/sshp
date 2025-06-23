package application

import intrastructure "github.com/mklbravo/sshp/infrastructure"

func ConnectToSSHHost(user, password, host string) {
	sshSession, _ := intrastructure.StartSSHSession(user, password, host)
	intrastructure.RunSSHShell(sshSession)
	defer sshSession.Close()
}
