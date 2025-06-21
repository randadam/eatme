#!/bin/bash
set -e

SCRIPT_DIR=$(dirname "$(realpath "$0")")
source "$SCRIPT_DIR/ml-venv.sh"

case "$1" in
    "install")
        ensure_deps
        ;;
    "add")
        if [ -z "$2" ]; then
            echo "Usage: eatme ml add <package_name> [--dev]"
            exit 1
        fi
        ensure_venv
        if [ "$3" == "--dev" ]; then
            echo "📦 Adding dev dependency: $2"
            . $VENV_DIR/bin/activate && uv pip install "$2"
            echo "⚠️  Don't forget to add $2 to pyproject.toml under [project.optional-dependencies].dev!"
        else
            echo "📦 Adding dependency: $2"
            . $VENV_DIR/bin/activate && uv pip install "$2"
            echo "⚠️  Don't forget to add $2 to pyproject.toml dependencies!"
        fi
        ;;
    "format")
        ensure_deps
        echo "🎨 Formatting code..."
        . $VENV_DIR/bin/activate && black ml-gateway/
        ;;
    "lint")
        ensure_deps
        echo "🔍 Linting code..."
        . $VENV_DIR/bin/activate && ruff check ml-gateway/
        ;;
    "test")
        ensure_deps
        echo "🧪 Running tests..."
        . $VENV_DIR/bin/activate && pytest ml-gateway/
        ;;
    *)
        echo "Usage: eatme ml [install|add|format|lint|test]"
        exit 1
        ;;
esac
