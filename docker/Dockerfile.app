FROM golang:1.23-alpine AS builder

WORKDIR /app

# Install build dependencies
RUN apk add --no-cache git make gcc musl-dev

# Copy entire project
COPY . .

# Download and verify dependencies
RUN go mod download
RUN go mod verify

# Build the application
RUN CGO_ENABLED=1 GOOS=linux go build -o server ./cmd/server/main.go

# Final stage
FROM alpine:latest

WORKDIR /app

# Install postgresql-client for pg_isready
RUN apk add --no-cache postgresql-client bash

# Copy the built binary
COPY --from=builder /app/server /app/server

# Copy migration script
COPY scripts/postgres-migrate.sh /scripts/postgres-migrate.sh
RUN chmod +x /scripts/postgres-migrate.sh

# Expose ports for gRPC, REST Gateway, and metrics
EXPOSE 8080 9008 8081

# Default command
CMD ["/app/server"]
