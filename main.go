package main

import (
	"fmt"
	"log"

	"github.com/mklbravo/sshp/infrastructure/sqlite"
	"github.com/mklbravo/sshp/internal/config"
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

	db.Ping()

	fmt.Println("DB initialized successfully at", cfg.DBPath)
}
