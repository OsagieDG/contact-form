# Use an official Golang runtime as a parent image
FROM golang:1.20.3-alpine

# Install air
RUN go install github.com/air-verse/air@v1.52.3

# Set the working directory to /app
WORKDIR /app

# Copy go.mod and go.sum files to the container
COPY go.mod ./

# Download the dependencies
RUN go mod download

# Copy the rest of the application code
COPY . .

# Build the Go application
RUN go build -o bin/api ./cmd/form

# Set the entry point of the container to air
CMD ["sh", "-c", "source .env && air -c .air.toml"]
