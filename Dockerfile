# Use the official Golang image as a base for building the app
FROM golang:1.22-alpine AS build

# Set the working directory inside the container
WORKDIR /app

# Install necessary tools for protobuf compilation and make
RUN apk add --no-cache make protobuf-dev bash

# Copy the Go modules and download dependencies early to cache these layers
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the project into the container
COPY . .

# Run the full pipeline defined in the Makefile
# Ensure PATH includes GOPATH/bin for protoc tools
RUN PATH=$PATH:$(go env GOPATH)/bin make all

# Use a lightweight container for the runtime environment
FROM alpine:latest

# Set the working directory inside the container
WORKDIR /app

# Copy only the necessary files from the build stage
COPY --from=build /app/main /app/main
COPY --from=build /app/.env /app/.env

# Expose the port your app runs on
EXPOSE 8000

# Run the Go application
CMD ["/app/main"]
