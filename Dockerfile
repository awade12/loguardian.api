# syntax=docker/dockerfile:1

# Stage 1: Build the Go application
FROM golang:latest AS build

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files to the workspace
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source code into the container
COPY . .

# Build the Go app
RUN CGO_ENABLED=0 GOOS=linux go build -o /go-translation-api

# Stage 2: Create the final image
FROM alpine:latest

# Set the Current Working Directory inside the container
WORKDIR /root/

# Copy the pre-built binary and .env file
COPY --from=build /go-translation-api .
COPY .env .

# Expose port 65000 to the outside world
EXPOSE 65000

# Command to run the executable
CMD ["./go-translation-api"]
