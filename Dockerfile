# Build stage
FROM golang:1.17-alpine AS builder

# Set the working directory in the container
WORKDIR /app

# Copy the go.mod and go.sum files to the working directory
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the entire source code from the current directory to the working directory
COPY . .

# Build the Go app
RUN go build -o hexa-notification .

# Final stage
FROM alpine:latest

# Set the working directory in the container
WORKDIR /app

# Copy the binary from the builder stage
COPY --from=builder /app/hexa-notification /app/hexa-notification

# Expose the port that the app will run on
EXPOSE 8080

# Run the binary executable
CMD ["/app/hexa-notification"]