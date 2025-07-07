#!/bin/bash
set -e

# Find the script's directory
SCRIPT_DIR=$(dirname "$(realpath "$0")")
API_DIR="$SCRIPT_DIR/../api"

echo "🧹 Cleaning up old API documentation..."
cd "$API_DIR"
rm -rf docs

echo "📝 Generating API documentation..."
swag init -g cmd/api/main.go
echo "✅ API documentation generated successfully!"
