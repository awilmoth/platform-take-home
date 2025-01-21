#!/bin/bash

# Exit on any error
set -e

# Function to log messages
log() {
    echo "[$(date +'%Y-%m-%d %H:%M:%S')] $*"
}

# Validate required environment variables
required_vars=("POSTGRES_HOST" "POSTGRES_PORT" "POSTGRES_USER" "POSTGRES_PASSWORD" "POSTGRES_DB")
for var in "${required_vars[@]}"; do
    if [ -z "${!var}" ]; then
        log "Error: $var is not set"
        exit 1
    fi
done

# Wait for Postgres to be ready
max_attempts=30
attempt=0

log "Waiting for Postgres to be ready..."
while ! pg_isready -h "$POSTGRES_HOST" -p "$POSTGRES_PORT" -U "$POSTGRES_USER"; do
    attempt=$((attempt + 1))
    if [ $attempt -ge $max_attempts ]; then
        log "Error: Postgres did not become ready in time"
        exit 1
    fi
    log "Waiting for Postgres (attempt $attempt/$max_attempts)..."
    sleep 5
done

# Construct Postgres DSN
export POSTGRES_DSN="postgres://${POSTGRES_USER}:${POSTGRES_PASSWORD}@${POSTGRES_HOST}:${POSTGRES_PORT}/${POSTGRES_DB}?sslmode=disable"

# Comprehensive migration diagnostics
log "Starting database migrations..."
log "Postgres Connection Details:"
log "Host: ${POSTGRES_HOST}"
log "Port: ${POSTGRES_PORT}"
log "Database: ${POSTGRES_DB}"
log "User: ${POSTGRES_USER}"

# Attempt to connect to the database
log "Verifying database connection..."
if PGPASSWORD="${POSTGRES_PASSWORD}" psql -h "${POSTGRES_HOST}" -p "${POSTGRES_PORT}" -U "${POSTGRES_USER}" -d "${POSTGRES_DB}" -c "\l" > /dev/null 2>&1; then
    log "Database connection successful"
else
    log "Error: Unable to connect to the database"
    log "Connection details:"
    PGPASSWORD="${POSTGRES_PASSWORD}" psql -h "${POSTGRES_HOST}" -p "${POSTGRES_PORT}" -U "${POSTGRES_USER}" -d "${POSTGRES_DB}" -c "\l" || true
    exit 1
fi

# Run migrations using a custom Go script to avoid OpenTelemetry context issues
log "Executing migration script..."
MIGRATION_SCRIPT=$(cat << 'EOF'
package main

import (
    "fmt"
    "os"
    "github.com/skip-mev/platform-take-home/store"
)

func main() {
    dsn := os.Getenv("POSTGRES_DSN")
    
    dbStore, err := store.NewPostgresBackedStore(dsn)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Failed to create database store: %v\n", err)
        os.Exit(1)
    }

    if err := dbStore.Migrate(); err != nil {
        fmt.Fprintf(os.Stderr, "Migration failed: %v\n", err)
        os.Exit(1)
    }

    fmt.Println("Migrations completed successfully")
}
EOF
)

MIGRATION_OUTPUT=$(go run <(echo "$MIGRATION_SCRIPT") 2>&1)
MIGRATE_EXIT_CODE=$?

if [ $MIGRATE_EXIT_CODE -eq 0 ]; then
    log "Migrations completed successfully"
    echo "$MIGRATION_OUTPUT"
else
    log "Error: Migration failed with exit code $MIGRATE_EXIT_CODE"
    log "Migration output:"
    echo "$MIGRATION_OUTPUT"
    
    # Additional diagnostic information
    log "Database table list:"
    PGPASSWORD="${POSTGRES_PASSWORD}" psql -h "${POSTGRES_HOST}" -p "${POSTGRES_PORT}" -U "${POSTGRES_USER}" -d "${POSTGRES_DB}" -c "\dt" || true
    
    exit $MIGRATE_EXIT_CODE
fi
