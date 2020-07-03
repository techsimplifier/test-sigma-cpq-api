ARG GO_VERSION=1.14

FROM golang:${GO_VERSION}-alpine AS builder

# Git is required for fetching the dependencies.
RUN apk add --no-cache ca-certificates git

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
COPY --from=builder /app/templates /templates

# Import the root ca-certificates (required for Let's Encrypt)
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

# Use non root user
USER 1000

# Serve
ENTRYPOINT ["test-sigma-cpq-api"]
CMD ["help"]