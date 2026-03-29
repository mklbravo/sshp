package main

import (
	"os"

	"github.com/mklbravo/sshp/cmd"
)

func main() {
	cmd.RootCmd.AddCommand(cmd.NewVersionCommand())
	cmd.RootCmd.AddCommand(cmd.NewAddCommand())

	if err := cmd.RootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
