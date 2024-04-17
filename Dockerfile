# Use the official Golang image as the base image
FROM golang:latest

# Set environment variables
ENV JWT_SECRET="Sintol"
ENV PORT="127.0.0.1:8080"

# Set the working directory inside the container
WORKDIR /app

# Copy the Go modules files
COPY go.mod .
COPY go.sum .

# Download dependencies
RUN go mod download

# Copy the rest of the application source code
COPY . .

# Build the Go application
RUN go build -o main .

# Expose the port the application runs on
EXPOSE 8080

# Command to run the executable
CMD ["./main"]