# Start with the Go base image
FROM golang:1.21 as builder

WORKDIR /app

# Copy the go.mod and go.sum files first and download the dependencies
# This step is cached unless go.mod or go.sum change
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the project files
COPY . .

# Build the Go app
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

# Use an Alpine base image for the final stage
FROM alpine:latest

# Install the ca-certificates package for HTTPS requests
RUN apk --no-cache add ca-certificates

# Set the working directory in the image
WORKDIR /root/

# Copy the built binary from the builder stage
COPY --from=builder /app/main .

# Copy the config.yaml file from the builder stage
COPY --from=builder /app/config.yaml .

# Command to run the executable
CMD ["./main"]
