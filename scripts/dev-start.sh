#!/bin/bash
set -e

echo "🚀 Starting dev environment..."

# Create .env file if it doesn't exist
if [ ! -f .env ]; then
  echo "📄 Creating default .env file..."
  cat <<EOF > .env
# Example environment variables
TURSO_URL=your-turso-url
ML_GATEWAY_URL=http://ml-gateway:8000
OLLAMA_URL=http://ollama:11434
EOF
fi

# Build and start services
echo "🐳 Starting Docker services..."
docker compose up --build -d

echo "✅ Dev environment is up and running!"
echo "🌐 Go API running at: http://localhost:8080"
echo "🧠 ML Gateway at: http://localhost:8000"
echo "🤖 Ollama API at: http://localhost:11434"
echo "📦 Chroma DB at: http://localhost:8002"
echo "📈 OTEL Collector at: http://localhost:4317"
echo "🖥️  Frontend: Run manually in separate terminal:"
echo "     cd app && pnpm dev"