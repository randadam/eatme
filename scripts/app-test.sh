#!/bin/bash
set -e

# Get the directory containing this script
SCRIPT_DIR=$(dirname "$(realpath "$0")")
APP_DIR="$SCRIPT_DIR/../app"

APP_PID=""
APP_PGID=""
cleanup() {
    echo "Cleaning up..."
    $SCRIPT_DIR/dev-stop.sh || true
    if [ -n "$APP_PGID" ]; then
        echo "Killing app process group: $APP_PGID"
        kill -TERM -"$APP_PGID" 2>/dev/null || true
        wait "$APP_PID" 2>/dev/null || true
        echo "Dev app stopped."
    fi
    echo "Done."
}

trap cleanup EXIT

# Start dev server
echo "Starting dev server..."
$SCRIPT_DIR/dev-start.sh

# Wait for dev server to start
tries=0
MAX_TRIES="${MAX_TRIES:-30}"
while ! curl -s http://localhost:8080/health > /dev/null; do
    sleep 1
    tries=$((tries+1))
    if [ $tries -ge $MAX_TRIES ]; then
        echo "Dev server failed to start after $MAX_TRIES seconds."
        exit 1
    fi
done

# Start dev app
echo "Starting dev app..."
$SCRIPT_DIR/app-start.sh &
APP_PID=$!
APP_PGID=$(ps -o pgid= $APP_PID | tr -d ' ')

# Run app tests
echo "Running tests..."
cd "$APP_DIR" && pnpm test:e2e
TEST_EXIT_CODE=$?

exit $TEST_EXIT_CODE
