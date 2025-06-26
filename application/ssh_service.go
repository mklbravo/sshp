package application

import (
    "github.com/mklbravo/sshp/domain"
    "github.com/mklbravo/sshp/infrastructure"
)

func ConnectToSSHHost(host domain.Host) error {
    sshSession, err := infrastructure.StartSSHSession(&host)
    if err != nil {
        return err
    }
    defer sshSession.Close()
    infrastructure.RunSSHShell(sshSession)
    return nil
}
