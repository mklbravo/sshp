#!/bin/bash

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
    # TODO: Add install dir to $PATH if is not already there
    echo "!!! Do not forget to add '${INSTALL_DIR}' to your \$PATH !!!"
else
    echo "Error: Failed to download '${APP_NAME}'"
    exit 1
fi
