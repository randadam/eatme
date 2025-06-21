# EatMe CLI Tool

The `eatme` command-line tool is your one-stop shop for all development tasks in the EatMe project. It provides a unified interface for managing services, development environments, and common development tasks.

## Global Commands

### Development Environment

```bash
# Start all services
eatme dev start

# Stop all services
eatme dev stop

## ML Gateway Commands

The `ml` subcommand provides tools for working with the ML Gateway service:

### Dependency Management

```bash
# Install all dependencies
eatme ml install

# Add a new dependency
eatme ml add <package_name>

# Add a development dependency
eatme ml add <package_name> --dev
```

### Development Tasks

```bash
# Format code with black
eatme ml format

# Lint code with ruff
eatme ml lint

# Run tests
eatme ml test
```

## API Commands

The `api` subcommand provides tools for working with the API service:

### Development Tasks

```bash
# Run API tests
eatme api test
```

## Virtual Environments

The CLI automatically manages virtual environments for you. You don't need to manually activate or deactivate them. Each service that requires a virtual environment (like the ML Gateway) will have its environment automatically created and managed by the CLI tool.

## Cleanup

```bash
# Clean up resources
eatme cleanup
```
