#!/bin/sh

# Default values
MIGRATE_CMD="up"

# Parse command line arguments
while [ "$#" -gt 0 ]; do
  case "$1" in
    up|down|version|force|goto) MIGRATE_CMD="$1"; shift ;;
    *) echo "Unknown command: $1"; exit 1 ;;
  esac
done

# Build the database URL from environment variables
DB_URL="postgres://${DB_USER}:${DB_PASSWORD}@${DB_HOST}:${DB_PORT}/${DB_NAME}?sslmode=disable"

# Execute migration command
case "$MIGRATE_CMD" in
  up)
    echo "Running migrations up..."
    migrate -path /go/src/github.com/hyorimitsu/knowledge-hub/backend/internal/infrastructure/persistence/migrations -database "$DB_URL" up
    ;;
  down)
    echo "Running migrations down..."
    migrate -path /go/src/github.com/hyorimitsu/knowledge-hub/backend/internal/infrastructure/persistence/migrations -database "$DB_URL" down
    ;;
  version)
    echo "Checking migration version..."
    migrate -path /go/src/github.com/hyorimitsu/knowledge-hub/backend/internal/infrastructure/persistence/migrations -database "$DB_URL" version
    ;;
  force)
    echo "Force setting migration version..."
    migrate -path /go/src/github.com/hyorimitsu/knowledge-hub/backend/internal/infrastructure/persistence/migrations -database "$DB_URL" force 1
    ;;
  goto)
    echo "Going to specific version..."
    migrate -path /go/src/github.com/hyorimitsu/knowledge-hub/backend/internal/infrastructure/persistence/migrations -database "$DB_URL" goto 1
    ;;
esac