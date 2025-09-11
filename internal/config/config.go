package config

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/joho/godotenv"
)

const APP_MODE_PROD = "PROD"
const APP_MODE_DEV = "DEV"

const ENV_KEY_MODE = "MODE"

type Config struct {
	Mode   string
	DBPath string
}

func Load() (*Config, error) {
	configDir, err := getConfigDir()
	if err != nil {
		return nil, err
	}

	godotenv.Load(filepath.Join(configDir, ".env"))

	mode, exists := os.LookupEnv(ENV_KEY_MODE)
	if !exists {
		mode = APP_MODE_PROD
	}

	mode = strings.ToUpper(mode)

	return &Config{
		Mode:   mode,
		DBPath: getDBPath(mode, configDir),
	}, nil
}

func getConfigDir() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	configDir := filepath.Join(home, ".config", "sshp")
	if err := os.MkdirAll(configDir, 0700); err != nil {
		return "", err
	}
	return configDir, nil
}

func getDBPath(mode string, configDir string) string {
	if mode == APP_MODE_DEV {
		return ".dev-data.db"
	}

	return filepath.Join(configDir, "data.db")

}
