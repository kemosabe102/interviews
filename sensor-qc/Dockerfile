# Dockerfile reference: https://github.com/GoogleCloudPlatform/golang-samples/blob/main/run/helloworld/Dockerfile
# Use the offical golang image to create a binary.
# This is based on Debian and sets the GOPATH to /go.
# https://hub.docker.com/_/golang
FROM golang:1.17-bullseye AS builder

# Create and change to the app directory.
WORKDIR /app

# Retrieve application dependencies.
# This allows the container build to reuse cached dependencies.
# Expecting to copy go.mod and if present go.sum.
COPY go.* ./
RUN go mod download

# Copy local code to the container image.
COPY . ./

# Build the binary.
RUN GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o sensor-qc

# Use the official Debian slim image for a lean production container.
# https://hub.docker.com/_/debian
FROM debian:bullseye-20211201-slim
RUN set -x && apt-get update && DEBIAN_FRONTEND=noninteractive apt-get install -y \
    ca-certificates && \
    rm -rf /var/lib/apt/lists/*

# Create and change to the app directory.
WORKDIR /app

# Copy the binary and config file to the production image from the builder stage.
COPY --from=builder /app/sensor-qc /app/sensor-qc
COPY --from=builder /app/config.yml /app/config.yml

CMD ["/app/sensor-qc"]