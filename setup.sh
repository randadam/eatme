#!/bin/bash
set -e

echo "🔧 Setting up EatMe CLI..."

# --- Installation Logic ---

# Check for Docker and Docker Compose
echo "🔎 Checking for Docker and Docker Compose..."
if ! command -v docker &> /dev/null; then
    echo "❌ Docker is not installed. Please install Docker and try again."
    echo "See: https://docs.docker.com/engine/install/"
    exit 1
fi

if ! docker info >/dev/null 2>&1; then
  echo "❌ Docker daemon is not running. Please start Docker and retry."
  exit 1
fi

if ! docker compose version >/dev/null 2>&1; then
  echo "❌ 'docker compose' command not found or not working."
  echo "Please ensure you have Docker Desktop or Docker Engine with the Compose V2 plugin installed."
  echo "See: https://docs.docker.com/compose/install/"
  exit 1
fi
echo "✅ Docker and Docker Compose are ready."

# Check for Node.js
echo "🔎 Checking for Node.js..."
if ! command -v node &> /dev/null; then
    echo "❌ Node.js is not installed. Please install Node.js and try again."
    echo "See: https://nodejs.org/"
    exit 1
fi

# Check Node.js version (we want at least 20)
NODE_VERSION=$(node -v | cut -d'v' -f2)
if [ "$(printf '%s\n' "20.0.0" "$NODE_VERSION" | sort -V | head -n1)" != "20.0.0" ]; then
    echo "❌ Node.js version must be 20 or higher. Found version: $NODE_VERSION"
    echo "Please upgrade Node.js and try again."
    exit 1
fi
echo "✅ Node.js $NODE_VERSION is ready."

# Check for pnpm
echo "🔎 Checking for pnpm..."
if ! command -v pnpm &> /dev/null; then
    echo "📦 pnpm not found. Installing..."
    npm install -g pnpm
    if [ $? -ne 0 ]; then
        echo "❌ Failed to install pnpm. Please try installing it manually:"
        echo "npm install -g pnpm"
        exit 1
    fi
else
    echo "✅ pnpm is already installed."
fi

# Check for Python
echo "🔎 Checking for Python..."
if ! command -v python3 &> /dev/null; then
    echo "❌ Python 3 is not installed. Please install Python 3 and try again."
    echo "See: https://www.python.org/downloads/"
    exit 1
fi

# Check Python version (we want at least 3.8)
PYTHON_VERSION=$(python3 -c 'import sys; print(".".join(map(str, sys.version_info[:2])))')
if [ "$(printf '%s\n' "3.8" "$PYTHON_VERSION" | sort -V | head -n1)" != "3.8" ]; then
    echo "❌ Python version must be 3.8 or higher. Found version: $PYTHON_VERSION"
    echo "Please upgrade Python and try again."
    exit 1
fi
echo "✅ Python $PYTHON_VERSION is ready."

# Install uv if missing
echo "🔎 Checking for uv package manager..."
if ! command -v uv &> /dev/null; then
  echo "📦 uv not found. Installing..."
  curl -LsSf https://github.com/astral-sh/uv/releases/latest/download/uv-installer.sh | sh
else
  echo "✅ uv is already installed."
fi

# Install Go if missing
echo "🔍 Checking for Go installation..."
if ! command -v go &> /dev/null; then
  echo "❌ Go is not installed. Please install Go (https://go.dev/doc/install) before proceeding."
  exit 1
fi
echo "✅ Go is installed."

# Install swag if missing
echo "🔍 Checking for swag CLI..."
if ! command -v swag &> /dev/null; then
  echo "📦 Installing swag CLI..."
  go install github.com/swaggo/swag/cmd/swag@latest

  # Add Go bin to PATH if necessary
  if [ "$(go env GOPATH)/bin" != "$(echo $PATH | tr ':' '\n' | grep "$(go env GOPATH)/bin")" ]; then
    export PATH=$PATH:$(go env GOPATH)/bin
    echo "Added $(go env GOPATH)/bin to PATH."
  fi
else
  echo "✅ swag CLI already installed."
fi

# --- PATH Setup ---

# Get the absolute path to the 'eatme' script
SCRIPT_DIR=$(dirname "$(realpath "$0")")
EATME_SCRIPT_PATH="$SCRIPT_DIR/scripts/eatme.sh"
INSTALL_PATH="/usr/local/bin/eatme"

echo "🔗 Creating symlink for 'eatme' command at $INSTALL_PATH..."

# Check for sudo permissions to create the symlink
if [ "$(id -u)" -ne 0 ]; then
  echo "This script requires sudo permissions to create a symlink in /usr/local/bin."
  echo "Please run with 'sudo ./setup.sh' or enter your password."
  sudo ln -sf "$EATME_SCRIPT_PATH" "$INSTALL_PATH"
else
  ln -sf "$EATME_SCRIPT_PATH" "$INSTALL_PATH"
fi

# Ensure all scripts are executable
chmod +x "$SCRIPT_DIR/scripts"/*.sh

echo "✅ Setup complete!"
echo "You can now use the 'eatme' command from anywhere."
