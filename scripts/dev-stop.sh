#!/bin/bash
set -e

echo "🛑 Stopping dev environment..."

# Stop Docker containers, removing volumes
docker compose down -v

echo "✅ Dev environment stopped."
