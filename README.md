<p align="center">
    <img src="https://img.shields.io/github/v/release/mklbravo/sshp?style=for-the-badge&logo=github&logoColor=%23cdd6f4&labelColor=%2345475a&color=%23cba6f7">
    <img src="https://img.shields.io/github/go-mod/go-version/mklbravo/sshp?style=for-the-badge&logo=go&logoColor=%23cdd6f4&labelColor=%2345475a&color=%2389b4fa">
    <img src="https://img.shields.io/github/actions/workflow/status/mklbravo/sshp/release.yml?style=for-the-badge&labelColor=%2345475a&color=%23a6e3a1">
</p>

# SSHP – Secure Shell Profiles

A modern Terminal User Interface (TUI) and CLI utility for managing SSH connection profiles quickly and interactively.

## Features
- Browse, search, and launch SSH connections with an intuitive terminal UI
- Manage host entries from a simple config file
- Fuzzy searching for hosts

## Installation

### Install via shell script (Recommended)
This is the preferred way to install the latest released version:

```sh
sh -c "$(curl -fsSL https://raw.githubusercontent.com/mklbravo/sshp/main/tools/install.sh)"
```

> **Note:** The script installs the `sshp` binary to `$HOME/.local/bin`. Make sure this directory is in your `$PATH` to use `sshp` globally.

### Build and install locally (Alternative)
1. Clone this repository:
   ```sh
   git clone <repo-url>
   cd sshp
   ```
2. Install:
   ```sh
   make install
   # Or directly: go install ./...
   ```

## Configuration

SSHP reads SSH host profiles from a JSON file. By default, it expects to find this file at:

```
$HOME/.config/sshp/hosts.json
```

### Example `hosts.json`
```json
[
  {
    "name": "production-server",
    "user": "alice",
    "address": "prod.example.com",
    "port": 22,
    "details": ["web", "production"]
  },
  {
    "name": "dev-box",
    "user": "devuser",
    "address": "10.0.0.5",
    "port": 2222,
    "details": []
  }
]
```
- `name`: label for the host (shown in TUI/fuzzy search)
- `address`: SSH address or IP
- `user`: SSH username
- `port`: (optional, defaults to 22)
- `details`: (optional, profile details)

## Usage
To launch the application, use:
```sh
sshp           # Recommended: launches the TUI
# Or, for development/testing:
go run main.go # From project root
```
- Use keyboard navigation and fuzzy search to manage SSH profiles.
- Make sure you have a valid `hosts.json` as described above.

## Development Environment (Optional)
- For isolated, reproducible dev environments (using Docker Compose):
  ```sh
  make devenv-build
  make devenv-start
  ```
  *(See Makefile for additional details.)*

## Troubleshooting
- **Missing hosts file**: Ensure your hosts JSON file is present at `$HOME/.config/sshp/hosts.json` or specify its location with an argument.

