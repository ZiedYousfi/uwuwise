# Build stage
FROM golang:alpine AS builder

# Install build dependencies
RUN apk add --no-cache git

WORKDIR /app

# Copy dependency files
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY . .

# Build the application
# -ldflags="-s -w" removes symbol table and debug information to reduce binary size
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o uwuwise .

# Run stage
FROM alpine:latest

# Install CA certificates for Discord API interaction via HTTPS
RUN apk add --no-cache ca-certificates

WORKDIR /root/

# Copy the binary from the builder stage
COPY --from=builder /app/uwuwise .

# Use an unprivileged user for security
RUN adduser -D uwuuser
USER uwuuser

# Command to run
ENTRYPOINT ["./uwuwise"]
