package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var version = "dev"

func NewVersionCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "version",
		Short: "Print the version number of SSHP",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(version)
		},
	}
}
