# Start from the official Golang base image
FROM golang:1.24.12-alpine AS builder

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy `go.mod` and `go.sum` files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the source code into the container
COPY . .

# Build the Go app
RUN go build -o main ./cmd

# Expose port 8080 (you can change it as per your application's port)
EXPOSE 8080

# Run the executable
CMD ["./main"]