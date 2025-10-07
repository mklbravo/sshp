package main

import (
	"log"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/mklbravo/sshp/application"
	"github.com/mklbravo/sshp/infrastructure/json"
	"github.com/mklbravo/sshp/infrastructure/ssh"
	"github.com/mklbravo/sshp/internal/config"
	"github.com/mklbravo/sshp/tui"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	hostRepository, err := json.NewJsonHostRepository(cfg.DataFilePath)
	if err != nil {
		log.Fatalf("Failed to load hosts: %v", err)
		os.Exit(1)
	}

	hostListUC := application.NewHostListUseCase(hostRepository)

	hostListView := tui.NewHostListView(hostListUC)

	tuiProgram := tea.NewProgram(hostListView)

	teaModel, err := tuiProgram.Run()

	if err != nil {
		log.Fatalf("Error running program: %v", err)
		os.Exit(1)
	}

	model := teaModel.(tui.Model)

	selectedHost := model.GetSelectedHost()
	if selectedHost == nil {
		log.Printf("No host selected")
		os.Exit(0)
	}

	hostConnectionUC := application.NewHostConnectionUseCase(
		ssh.NewSSHConnectionService(),
	)

	err = hostConnectionUC.Execute(selectedHost)
	if err != nil {
		log.Fatalf("Failed to connect to SSH host: %v", err)
		os.Exit(1)
	}
}
