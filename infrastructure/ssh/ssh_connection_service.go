package ssh

import (
	"errors"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"
	"time"

	"github.com/mklbravo/sshp/domain/entity"
	"golang.org/x/crypto/ssh"
	"golang.org/x/crypto/ssh/agent"
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
	auth, closeAgent := newPublicKeyAuthentication()
	defer closeAgent()

	// SSH client configuration
	config := &ssh.ClientConfig{
		User: profile.Username.GetValue(),
		Auth: []ssh.AuthMethod{
			auth,
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

func newPublicKeyAuthentication() (ssh.AuthMethod, func()) {
	var agentConnection net.Conn
	agentAttempted := false
	filesAttempted := false

	signers := func() ([]ssh.Signer, error) {
		if !agentAttempted {
			agentAttempted = true
			if socket := os.Getenv("SSH_AUTH_SOCK"); socket != "" {
				connection, err := net.DialTimeout("unix", socket, 5*time.Second)
				if err == nil {
					agentConnection = connection
					keys, err := agent.NewClient(connection).Signers()
					if err == nil && len(keys) > 0 {
						return keys, nil
					}
				}
			}
		}

		if filesAttempted {
			return nil, nil
		}
		filesAttempted = true
		return getPrivateKeySigners()
	}

	// Both sources use publickey authentication, so retry the same method
	// to load local keys only after the agent's keys have been rejected.
	auth := ssh.RetryableAuthMethod(ssh.PublicKeysCallback(signers), 2)
	return auth, func() {
		if agentConnection != nil {
			agentConnection.Close()
		}
	}
}

func getPrivateKeySigners() ([]ssh.Signer, error) {
	homeFolder, err := os.UserHomeDir()
	if err != nil {
		return nil, fmt.Errorf("find home directory for SSH keys: %w", err)
	}

	privateKeyCandidates := []string{
		"id_ed25519",
		"id_rsa",
		"id_ecdsa",
	}

	var signers []ssh.Signer
	var keyErrors []error
	for _, keyFile := range privateKeyCandidates {
		keyPath := filepath.Join(homeFolder, ".ssh", keyFile)
		key, err := os.ReadFile(keyPath)
		if err != nil {
			if !errors.Is(err, os.ErrNotExist) {
				keyErrors = append(keyErrors, fmt.Errorf("read private key %s: %w", keyPath, err))
			}
			continue
		}
		signer, err := ssh.ParsePrivateKey(key)
		if err != nil {
			keyErrors = append(keyErrors, fmt.Errorf("parse private key %s: %w", keyPath, err))
			continue
		}
		signers = append(signers, signer)
	}

	if len(signers) > 0 {
		return signers, nil
	}
	if len(keyErrors) > 0 {
		return nil, fmt.Errorf("load local SSH private keys: %w", errors.Join(keyErrors...))
	}
	return nil, errors.New("agent authentication unavailable or unsuccessful and no private key found in ~/.ssh/")
}
