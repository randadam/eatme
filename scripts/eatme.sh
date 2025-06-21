#!/bin/bash
set -e

# Find the script's directory to allow execution from any path
SCRIPT_DIR=$(dirname "$(realpath "$0")")

# Main command dispatcher
case "$1" in
  dev)
    case "$2" in
      start)
        "$SCRIPT_DIR/dev-start.sh"
        ;;
      stop)
        "$SCRIPT_DIR/dev-stop.sh"
        ;;
      *)
        echo "Usage: $0 dev [start|stop]"
        exit 1
        ;;
    esac
    ;;
  ml)
    shift # Remove 'ml' from the arguments
    "$SCRIPT_DIR/ml-dev.sh" "$@"
    ;;
  api)
    case "$2" in
      test)
        "$SCRIPT_DIR/api-test.sh"
        ;;
      docs)
        "$SCRIPT_DIR/api-docs.sh"
        ;;
      *)
        echo "Usage: $0 api [test|docs]"
        exit 1
        ;;
    esac
    ;;
  app)
    case "$2" in
      start)
        "$SCRIPT_DIR/app-start.sh"
        ;;
      test)
        "$SCRIPT_DIR/app-test.sh"
        ;;
      *)
        echo "Usage: $0 app [start|test]"
        exit 1
        ;;
    esac
    ;;
  cleanup)
    "$SCRIPT_DIR/cleanup.sh"
    ;;
  *)
    echo "Usage: $0 [dev|ml|api|cleanup]"
    echo
    echo "Commands:"
    echo "  dev <command>       - Manage development environment"
    echo "    start             - Start development environment"
    echo "    stop              - Stop development environment"
    echo "  ml <command>        - ML Gateway development tasks"
    echo "    install           - Install dependencies"
    echo "    add <pkg> [--dev] - Add a new dependency"
    echo "    format            - Format code with black"
    echo "    lint              - Lint code with ruff"
    echo "    test              - Run tests"
    echo "  api <command>       - API development commands"
    echo "    test              - Run API tests"
    echo "    docs              - Generate API documentation"
    echo "  app <command>       - App development commands"
    echo "    start             - Start app"
    echo "    test              - Run app tests"
    echo "  cleanup             - Clean up resources"
    exit 1
    ;;
esac
