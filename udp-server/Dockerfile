# Use the official Go 1.19 base image
FROM golang:1.19 as builder

# Set the working directory
WORKDIR /app

# Copy the Go source code into the container
COPY go.mod go.sum ./
COPY main.go .

# Build the Go binary
RUN go build -o udp-server main.go

# Use a lightweight base image for the final container
FROM debian:buster-slim

# Set the working directory
WORKDIR /app

# Copy the built binary from the builder stage
COPY --from=builder /app/udp-server .

# Expose the UDP port
EXPOSE 9098/udp

# Run the UDP server
CMD ["./udp-server"]
