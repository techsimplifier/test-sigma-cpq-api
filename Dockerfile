ARG GO_VERSION=1.14

FROM golang:${GO_VERSION} AS builder

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download -x

# Copy the source files from the current directory to the Working Directory inside the container
COPY . .

# Build the app
RUN CGO_ENABLED=0 GO111MODULE=on GOOS=linux GOARCH=amd64 go build -ldflags '-w -extldflags "-static"'

# Final build should only be scratch
FROM scratch

# Add maintainer info
LABEL maintainer="Tech Simplifier"

# Copy binary from build process to container
COPY --from=builder /app/test-sigma-cpq-api /usr/bin/test-sigma-cpq-api

# Use non root user
USER 1000

# Serve
ENTRYPOINT ["test-sigma-cpq-api"]
CMD ["help"]