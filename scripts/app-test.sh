#!/bin/bash
set -e

# Get the directory containing this script
SCRIPT_DIR=$(dirname "$(realpath "$0")")
APP_DIR="$SCRIPT_DIR/../app"

# Run all tests in the app directory and its subdirectories
cd "$APP_DIR" && pnpm test:unit