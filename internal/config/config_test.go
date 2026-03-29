package config

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetDataFilePath(t *testing.T) {
	configDir := "/home/user/.ssh/sshp"

	tests := []struct {
		name     string
		mode     string
		expected string
	}{
		{
			name:     "prod mode",
			mode:     APP_MODE_PROD,
			expected: filepath.Join(configDir, "hosts.json"),
		},
		{
			name:     "dev mode",
			mode:     APP_MODE_DEV,
			expected: "hosts.dev.json",
		},
		{
			name:     "unknown mode defaults to prod",
			mode:     "UNKNOWN",
			expected: filepath.Join(configDir, "hosts.json"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := getDataFilePath(tt.mode, configDir)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestLoad_ModeFromEnv(t *testing.T) {
	// Test with MODE=DEV
	t.Setenv(ENV_KEY_MODE, "dev")
	defer os.Unsetenv(ENV_KEY_MODE)

	config, err := Load()
	assert.NoError(t, err)
	assert.Equal(t, APP_MODE_DEV, config.Mode)
	assert.Equal(t, "hosts.dev.json", config.DataFilePath)
}

func TestLoad_DefaultMode(t *testing.T) {
	// Test without MODE set
	os.Unsetenv(ENV_KEY_MODE)

	config, err := Load()
	assert.NoError(t, err)
	assert.Equal(t, APP_MODE_PROD, config.Mode)
	assert.Contains(t, config.DataFilePath, "hosts.json")
}

func TestLoad_ModeUppercase(t *testing.T) {
	// Test with lowercase mode
	t.Setenv(ENV_KEY_MODE, "dev")
	defer os.Unsetenv(ENV_KEY_MODE)

	config, err := Load()
	assert.NoError(t, err)
	assert.Equal(t, APP_MODE_DEV, config.Mode)
}

func TestGetConfigDir(t *testing.T) {
	// Test that it returns a path under home
	dir, err := getConfigDir()
	assert.NoError(t, err)
	assert.Contains(t, dir, ".ssh")
	assert.Contains(t, dir, "sshp")

	// Verify the directory exists (it creates it)
	_, err = os.Stat(dir)
	assert.NoError(t, err)
}
