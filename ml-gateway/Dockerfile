# Development stage
FROM python:3.12-slim-bookworm
COPY --from=ghcr.io/astral-sh/uv:latest /uv /uvx /bin/

WORKDIR /app

# Copy project configuration
COPY pyproject.toml .

# Install dependencies
RUN uv pip install --system --no-cache -e ".[dev]"

# Copy source code
COPY . .

# Expose port
EXPOSE 8000

# Run with uvicorn in development mode (auto-reload)
CMD ["uvicorn", "app.main:app", "--host", "0.0.0.0", "--port", "8000", "--reload"]
