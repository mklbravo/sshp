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
	"github.com/spf13/cobra"
)

func main() {
	rootCmd := createRootCommand()
	if err := rootCmd.Execute(); err != nil {
		log.Fatalf("Error executing command: %v", err)
		os.Exit(1)
	}
}

func createRootCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "sshp",
		Short: "SSHP is a terminal-based SSH host manager and connector.",
		Run: func(cmd *cobra.Command, args []string) {
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
		},
	}
}
