#!/bin/bash

setup_path_if_needed() {
    if [[ ":$PATH:" != *":$INSTALL_DIR:"* ]]; then
        case "$SHELL" in
            */bash) config_file="$HOME/.bashrc" ;;
            */zsh) config_file="$HOME/.zshrc" ;;
            *) config_file="$HOME/.profile" ;;
        esac

        echo "export PATH=\"$INSTALL_DIR:\$PATH\"" >> "$config_file"
        echo "Added '$INSTALL_DIR' to your PATH in $config_file"
        echo "Restart your shell or run 'source $config_file' to use it immediately"
    else
        echo "'$INSTALL_DIR' is already in your PATH"
    fi
}

# Detect OS
OS=$(uname -s)
ARCH=$(uname -m)

# Normalize architecture names
case "$ARCH" in
x86_64) ARCH="amd64" ;;
aarch64 | arm64) ARCH="arm64" ;;
*) ;;
esac

# Check supported OS
if [[ "$OS" != "Linux" && "$OS" != "Darwin" ]]; then
    echo "Unsupported OS: $OS"
    exit 1
fi

# Check supported architecture
if [[ "$ARCH" != "amd64" && "$ARCH" != "arm64" ]]; then
    echo "Unsupported architecture: $ARCH"
    exit 1
fi

echo "Detected OS: $OS"
echo "Detected architecture: $ARCH"

API_URL="https://api.github.com/repos/mklbravo/sshp/releases/latest"
DOWNLOAD_URL=$(curl -s "$API_URL" | grep -i "browser_download_url.*${OS}.*${ARCH}" | head -1 | cut -d '"' -f 4)

if [[ -z "$DOWNLOAD_URL" ]]; then
    echo "Error: Could not find download URL for $OS $ARCH"
    exit 1
fi

# Create install directory if it doesn't exist
INSTALL_DIR="$HOME/.local/bin"
mkdir -p "$INSTALL_DIR"

APP_NAME="sshp"
APP_PATH="${INSTALL_DIR}/${APP_NAME}"

if curl -L "$DOWNLOAD_URL" -o "${APP_PATH}"; then
    chmod +x "${APP_PATH}"
    echo "'${APP_NAME}' installed in '${INSTALL_DIR}'"
    setup_path_if_needed
else
    echo "Error: Failed to download '${APP_NAME}'"
    exit 1
fi
