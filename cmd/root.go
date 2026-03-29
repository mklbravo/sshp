package cmd

import (
	"log"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/spf13/cobra"

	"github.com/mklbravo/sshp/application"
	"github.com/mklbravo/sshp/infrastructure/json"
	"github.com/mklbravo/sshp/infrastructure/ssh"
	"github.com/mklbravo/sshp/internal/config"
	"github.com/mklbravo/sshp/tui"
)

// RootCmd is the main CLI entrypoint (TUI and subcommands attach here)
var RootCmd = createRootCommand()

func createRootCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "sshp",
		Short: "SSHP is a terminal-based SSH host manager and connector.",
		Run: func(cmd *cobra.Command, args []string) {
			cfg, err := config.Load()
			if err != nil {
				log.Fatalf("Failed to load config: %v", err)
			}

			hostRepository, err := json.NewJsonProfileRepository(cfg.DataFilePath)
			if err != nil {
				log.Fatalf("Failed to load hosts: %v", err)
				os.Exit(1)
			}

			profileListUC := application.NewProfileListUseCase(hostRepository)
			profileListView := tui.NewProfileListView(profileListUC)
			tuiProgram := tea.NewProgram(profileListView)
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

			profileConnectionUC := application.NewProfileConnectionUseCase(
				ssh.NewSSHConnectionService(),
			)
			err = profileConnectionUC.Execute(selectedHost)
			if err != nil {
				log.Fatalf("Failed to connect to SSH host: %v", err)
				os.Exit(1)
			}
		},
	}
}
