#!/bin/bash
set -e

echo "🧹 Cleaning up resources..."

# Prompt to remove Ollama models
read -p "❓ Do you want to remove all downloaded Ollama models? (y/N): " confirm_ollama
if [[ "$confirm_ollama" =~ ^[Yy]$ ]]; then
  echo "🗑️ Removing Ollama models..."
  if [ -n "$(ollama list | awk 'NR>1')" ]; then
    ollama rm $(ollama list | awk 'NR>1 {print $1}')
  else
    echo "No Ollama models to remove."
  fi
fi

# Prompt to delete .env
read -p "❓ Do you want to delete the .env file? (y/N): " confirm_env
if [[ "$confirm_env" =~ ^[Yy]$ ]]; then
  echo "🧾 Deleting .env file..."
  rm -f .env
fi

echo "✅ Cleanup complete."
