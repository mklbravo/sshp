package cmd

import (
	"github.com/mklbravo/sshp/domain/entity"
	"github.com/mklbravo/sshp/infrastructure/json"
	"github.com/mklbravo/sshp/internal/config"
	"github.com/spf13/cobra"
)

func NewAddCommand() *cobra.Command {
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
