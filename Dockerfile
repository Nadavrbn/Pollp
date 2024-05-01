# Use the official golang image as base image
FROM golang:1.22-alpine as builder

# Set the current working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum to install dependencies
COPY go.mod go.sum ./

# Download and install Go module dependencies
RUN go mod download

# Copy the rest of the application source code
COPY . .

# Build the Go application
RUN go build -o main .

# Start a new stage from scratch
FROM alpine:latest

# Set the current working directory inside the container
WORKDIR /app

# Copy the built executable from the previous stage
COPY --from=builder /app/main .

# Expose the port on which the application will listen
EXPOSE 8080

# Command to run the executable
CMD ["./main"]