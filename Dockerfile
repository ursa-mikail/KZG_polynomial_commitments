# Use official Golang image as build environment
# FROM golang:1.22 AS builder
FROM golang:1.24 AS builder

# Set working directory
WORKDIR /app

# Copy go.mod and go.sum first (to leverage Docker layer cache)
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the source code
COPY . .

# Build the Go binary
RUN go build -o kzg-demo main.go

# Final minimal image
FROM debian:bookworm-slim

WORKDIR /app

# Copy only the binary from builder stage
COPY --from=builder /app/kzg-demo .

# Default command
ENTRYPOINT ["./kzg-demo"]

