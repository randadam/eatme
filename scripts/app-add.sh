#!/bin/bash
set -e

# Find the script's directory to allow execution from any path
SCRIPT_DIR=$(dirname "$(realpath "$0")")
APP_DIR="$SCRIPT_DIR/../app"

cd "$APP_DIR"

case "$1" in
    "dep")
        echo "Adding dependency..."
        pnpm add "$2"
        ;;
    "comp")
        echo "Adding component..."
        pnpm dlx shadcn@latest add "$2"
        ;;
    *)
        echo "Usage: eatme app add [dep|comp] <package_name>"
        exit 1
        ;;
esac