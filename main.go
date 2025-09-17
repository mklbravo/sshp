package main

import (
	"log"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/mklbravo/sshp/application"
	"github.com/mklbravo/sshp/infrastructure/sqlite"
	"github.com/mklbravo/sshp/internal/config"
	"github.com/mklbravo/sshp/tui"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	db, err := sqlite.GetDBConnection(cfg.DBPath)

	if err != nil {
		log.Fatalf("Failed to initialize db: %v", err)
	}

	hostListUC := application.NewHostListUseCase(sqlite.NewHostRepository(db))

	hostListView := tui.NewHostListView(hostListUC)

	p := tea.NewProgram(hostListView)

	if _, err := p.Run(); err != nil {
		log.Fatalf("Error running program: %v", err)
		os.Exit(1)
	}

}
