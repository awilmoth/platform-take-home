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

# Install required packages
RUN apk add --no-cache postgresql-client bash go

# Copy the built binary and necessary Go files
COPY --from=builder /app/server /app/server
COPY --from=builder /app/go.mod /app/go.mod
COPY --from=builder /app/go.sum /app/go.sum
COPY --from=builder /app/store /app/store

# Set up Go environment
ENV GOPATH=/go
ENV PATH=$PATH:/usr/local/go/bin:$GOPATH/bin

# Copy migration script
COPY scripts/postgres-migrate.sh /scripts/postgres-migrate.sh
RUN chmod +x /scripts/postgres-migrate.sh

# Expose ports for gRPC, REST Gateway, and metrics
EXPOSE 8080 9008 8081

# Default command
CMD ["/app/server"]
