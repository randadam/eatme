#!/bin/bash
set -e

# Get the directory containing this script
SCRIPT_DIR=$(dirname "$(realpath "$0")")
APP_DIR="$SCRIPT_DIR/../app"

# Run the app
cd "$APP_DIR" && pnpm dev
