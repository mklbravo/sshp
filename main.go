package main

import (
	"fmt"
	"log"

	"github.com/mklbravo/sshp/application"
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

	fmt.Printf("Using database at: %s\n", cfg.DBPath)
	hostListUC := application.NewHostListUseCase(sqlite.NewHostRepository(db))

	hosts, err := hostListUC.Execute()
	if err != nil {
		log.Fatalf("Failed to list hosts: %v", err)
	}
	fmt.Println("Listing all hosts:")
	for _, host := range hosts {
		fmt.Printf("Host: %s, IP: %s, Port: %d\n", host.Name, host.IP, host.Port)
	}
}
