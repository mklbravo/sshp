package main

import (
	"fmt"
	"log"

	"github.com/mklbravo/sshp/infrastructure"
	"github.com/mklbravo/sshp/internal/config"
)

func main() {
	cfg, err := config.LoadDev()
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}
	db, err := infrastructure.InitDB(cfg.DBPath)
	if err != nil {
		log.Fatalf("failed to initialize db: %v", err)
	}
	db.Ping()

	fmt.Println("DB initialized successfully at", cfg.DBPath)
}
