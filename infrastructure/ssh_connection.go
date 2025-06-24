package infrastructure

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/mklbravo/sshp/domain"
	"golang.org/x/crypto/ssh"
	"golang.org/x/term"
)

func StartSSHSession(host *domain.Host) (*ssh.Session, error) {

	// SSH client configuration
	config := &ssh.ClientConfig{
		User: host.Username,
		Auth: []ssh.AuthMethod{
			ssh.PublicKeys(getPrivateKeySigner()), // Use private key for authentication
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(), // WARNING: for testing only
	}

	// Connect to the SSH server
	client, err := ssh.Dial("tcp", fmt.Sprintf("%s:%d", host.IP, host.Port), config)
	if err != nil {
		return nil, err
	}

	// Open a session
	session, err := client.NewSession()
	if err != nil {
		client.Close()
		return nil, err
	}

	session.Stdin = os.Stdin
	session.Stdout = os.Stdout
	session.Stderr = os.Stderr

	return session, nil
}

func RunSSHShell(this *ssh.Session) {
	// Set up terminal modes
	modes := ssh.TerminalModes{
		ssh.ECHO:          1,     // enable echoing
		ssh.TTY_OP_ISPEED: 14400, // input speed = 14.4kbaud
		ssh.TTY_OP_OSPEED: 14400, // output speed = 14.4kbaud
	}

	// Get current terminal size
	fd := int(os.Stdin.Fd())
	width, height, err := term.GetSize(fd)
	if err != nil {
		width, height = 80, 24
	}

	// Request PTY
	if err := this.RequestPty("xterm", height, width, modes); err != nil {
		log.Fatalf("request for pseudo terminal failed: %s", err)
	}

	// Handle terminal resize
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGWINCH)
	go func() {
		for range signals {
			width, height, _ := term.GetSize(fd)
			this.WindowChange(height, width)
		}
	}()

	// Start shell
	if err := this.Shell(); err != nil {
		log.Fatalf("failed to start shell: %s", err)
	}
	this.Wait()
}

func getPrivateKeySigner() ssh.Signer {
	// TODO: Make this configurable
	keyPath := fmt.Sprintf("%s/.ssh/id_rsa", os.Getenv("HOME"))

	key, err := os.ReadFile(keyPath)
	if err != nil {
		log.Fatalf("unable to read private key: %v", err)
	}

	signer, err := ssh.ParsePrivateKey(key)
	if err != nil {
		log.Fatalf("unable to parse private key: %v", err)
	}

	return signer
}
