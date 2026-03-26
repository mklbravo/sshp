package main

import (
	"fmt"
	"log"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/mklbravo/sshp/application"
	"github.com/mklbravo/sshp/domain/entity"
	"github.com/mklbravo/sshp/infrastructure/json"
	"github.com/mklbravo/sshp/infrastructure/ssh"
	"github.com/mklbravo/sshp/internal/config"
	"github.com/mklbravo/sshp/tui"
	"github.com/spf13/cobra"
)

var version = "dev"

func main() {
	rootCmd := createRootCommand()
	rootCmd.AddCommand(createVersionCommand())
	rootCmd.AddCommand(createAddCommand())

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

func createVersionCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "version",
		Short: "Print the version number of SSHP",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(version)
		},
	}
}

func createAddCommand() *cobra.Command {
	var (
		name    string
		user    string
		address string
		port    int
		details []string
	)

	cmd := &cobra.Command{
		Use:   "add",
		Short: "Add a new SSH profile",
		RunE: func(cmd *cobra.Command, args []string) error {
			cfg, err := config.Load()
			if err != nil {
				return err
			}

			repo, err := json.NewJsonProfileRepository(cfg.DataFilePath)
			if err != nil {
				return err
			}

			profiles, err := repo.FindAll()
			if err != nil {
				return err
			}

			id := len(profiles)
			profile, err := entity.NewProfile(id, name, user, address, port, "default", details)
			if err != nil {
				return err
			}

			if err := repo.Save(profile); err != nil {
				return err
			}

			cmd.Printf("Profile added: %s\n", name)
			return nil
		},
	}
	cmd.Flags().StringVarP(&name, "name", "n", "", "Name of the profile (required)")
	cmd.Flags().StringVarP(&user, "user", "u", "", "SSH username (required)")
	cmd.Flags().StringVarP(&address, "address", "a", "", "SSH address or IP (required)")
	cmd.Flags().IntVarP(&port, "port", "p", 22, "SSH port")
	cmd.Flags().StringSliceVar(&details, "details", nil, "Details/tags for the profile (comma-separated or repeatable)")

	cmd.MarkFlagRequired("name")
	cmd.MarkFlagRequired("user")
	cmd.MarkFlagRequired("address")

	return cmd
}
