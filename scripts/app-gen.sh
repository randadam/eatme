#!/bin/bash
set -e

# Find the script's directory to allow execution from any path
SCRIPT_DIR=$(dirname "$(realpath "$0")")
APP_DIR="$SCRIPT_DIR/../app"

cd "$APP_DIR"

# Main command dispatcher
case "$1" in
    client)
        echo "Generating API docs..."
        $SCRIPT_DIR/api-docs.sh
        echo "Generating client..."
        pnpm gen:client
        echo "✅ API client generated successfully!"
        ;;
    *)
        echo "Usage: $0 app gen [client]"
        exit 1
        ;;
esac