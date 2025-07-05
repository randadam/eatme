# Setup Guide

This guide will help you get started with the EatMe development environment.

## Prerequisites

Before you begin, ensure you have the following installed on your system:

- **Docker:** The containerization platform.
- **Docker Compose (V2 Plugin):** The tool for defining and running multi-container Docker applications.
- **Python 3.8+:** Required for the ML Gateway service.
- **Node.js 20+:** Required for the frontend and API documentation.
  - The setup script will automatically install `pnpm` if it's not present.
- **Ollama:** Must be installed and running locally. The application expects Ollama to be available at `http://localhost:11434`.

## One-Time Setup

To get started, you need to run the one-time setup script. This script will:

1. Verify that Docker and Docker Compose are installed and running
2. Check for Python 3.8+ installation
3. Install development tools:
   - `uv` for Python package management
   - `pnpm` for Node.js package management
   - `swag` for API documentation
4. Create a globally accessible `eatme` command

Open your terminal in the project root and run:

```bash
sudo ./setup.sh
```

You will be prompted for your password because the script needs `sudo` permissions to create the symbolic link.

## Development Workflow

Once setup is complete, you can use the `eatme` CLI tool to manage your development environment. See the [CLI Documentation](cli/README.md) for detailed usage.

### Common Commands

```bash
# Start all backend services
eatme dev start

# Start frontend
eatme app start

# Stop all services
eatme dev stop

# Clean up resources
eatme cleanup
```

## Services

When the development environment is running, the following services will be available:

| Service           | URL                            | Purpose                               |
| ----------------  | ------------------------------ | ------------------------------------- |
| **Go API**        | `http://localhost:8080`        | Main application backend              |
| **ML Gateway**    | `http://localhost:8000`        | Gateway for machine learning tasks    |
| **Chroma DB**     | `http://localhost:8002`        | Vector database for embeddings        |
| **OTEL Collector**| `http://localhost:4317` (gRPC) | Collects observability data           |
| **Frontend**      | `http://localhost:5173`        | React application                     |
