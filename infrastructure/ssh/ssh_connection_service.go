package ssh

import (
	"log"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"

	"github.com/mklbravo/sshp/domain/entity"
	"golang.org/x/crypto/ssh"
	"golang.org/x/term"
)

type SSHConnectionService struct {
}

func NewSSHConnectionService() *SSHConnectionService {
	return &SSHConnectionService{}
}

func (this *SSHConnectionService) ConnectToHost(profile *entity.Profile) error {
	sshSession, err := StartSSHSession(profile)
	if err != nil {
		return err
	}
	defer sshSession.Close()
	RunSSHShell(sshSession)
	return nil
}

func StartSSHSession(profile *entity.Profile) (*ssh.Session, error) {

	// SSH client configuration
	config := &ssh.ClientConfig{
		User: profile.Username.GetValue(),
		Auth: []ssh.AuthMethod{
			ssh.PublicKeys(getPrivateKeySigner()), // Use private key for authentication
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(), // WARNING: for testing only
	}

	// Connect to the SSH server
	client, err := ssh.Dial("tcp", profile.GetFullAddress(), config)
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

	// Set terminal to raw mode
	oldState, err := term.MakeRaw(fd)
	if err != nil {
		log.Fatalf("failed to set terminal to raw mode: %s", err)
	}
	defer term.Restore(fd, oldState)

	// Handle terminal resize
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGWINCH)
	go func() {
		for range signals {
			width, height, _ := term.GetSize(fd)
			this.WindowChange(height, width)
		}
	}()
	// Send initial size
	this.WindowChange(height, width)

	// Forward SIGINT (Ctrl+C) to remote process
	// MakeRaw will output Ctrl+C directly to the remote session,
	// but in case we need to handle it differently, we set up this forwarding.
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)
	go func() {
		for range interrupt {
			_ = this.Signal(ssh.SIGINT)
		}
	}()

	// Start shell
	if err := this.Shell(); err != nil {
		log.Fatalf("failed to start shell: %s", err)
	}
	this.Wait()
}

func getPrivateKeySigner() ssh.Signer {
	homeFolder := os.Getenv("HOME")

	if homeFolder == "" {
		log.Fatalf("HOME environment variable is not set")
	}

	privateKeyCandidates := []string{
		"id_ed25519",
		"id_rsa",
		"id_ecdsa",
	}

	for _, keyFile := range privateKeyCandidates {

		keyPath := filepath.Join(homeFolder, ".ssh", keyFile)
		key, err := os.ReadFile(keyPath)
		if err != nil {
			// Try the next candidate if the file doesn't exist
			continue
		}
		signer, err := ssh.ParsePrivateKey(key)
		if err != nil {
			log.Fatalf("Unable to parse private key (%s): %v", keyPath, err)
		}

		return signer
	}

	log.Fatalf("No valid private key found in ~/.ssh/")
	return nil
}
