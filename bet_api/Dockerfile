FROM golang:1.23.2-alpine3.20 AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum for dependency management
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the rest of the source code
COPY . .

# Build the Go app
RUN go build -o bet-api

# Run stage: use a smaller base image
FROM alpine:latest

# Set the working directory
WORKDIR /root/

# Copy the built app from the builder stage
COPY --from=builder /app/bet-api .

# Expose the port your app runs on
EXPOSE 8002

# Command to run the Go app
CMD ["./bet-api"]
