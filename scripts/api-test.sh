#!/bin/bash
set -e

# Get the directory containing this script
SCRIPT_DIR=$(dirname "$(realpath "$0")")
API_DIR="$SCRIPT_DIR/../api"

# Run all tests in the api directory and its subdirectories
cd "$API_DIR" && go test -v ./...
