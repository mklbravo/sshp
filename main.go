package main


import (
	"github.com/mklbravo/sshp/application"
)

import (
    "log"
    "github.com/mklbravo/sshp/internal/config"
    "github.com/mklbravo/sshp/infrastructure"
)

func main() {
    cfg, err := config.Load()
    if err != nil {
        log.Fatalf("failed to load config: %v", err)
    }
    db, err := infrastructure.InitDB(cfg.DBPath)
    if err != nil {
        log.Fatalf("failed to initialize db: %v", err)
    }
    // DB initialized. Add further logic here if needed.
}

