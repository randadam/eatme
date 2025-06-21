#!/bin/bash
set -e

VENV_DIR="ml-gateway/.venv"
PYTHON="python3"

ensure_venv() {
    if [ ! -d "$VENV_DIR" ]; then
        echo "🔧 Creating Python virtual environment..."
        $PYTHON -m venv $VENV_DIR
        $VENV_DIR/bin/python -m pip install uv
    fi
}

ensure_deps() {
    ensure_venv
    echo "📦 Installing dependencies..."
    . $VENV_DIR/bin/activate && uv pip install -e "ml-gateway/[dev]"
}
