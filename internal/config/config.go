package config

import (
    "os"
    "path/filepath"
)

type Config struct {
    DBPath string
}

func Load() (*Config, error) {
    home, err := os.UserHomeDir()
    if err != nil {
        return nil, err
    }
    dbDir := filepath.Join(home, ".config", "sshp")
    if err := os.MkdirAll(dbDir, 0700); err != nil {
        return nil, err
    }
    dbPath := filepath.Join(dbDir, "data.db")
    return &Config{DBPath: dbPath}, nil
}
