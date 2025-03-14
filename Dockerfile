# Build stage
FROM golang:1.23.1-alpine AS builder

# Install necessary build tools
RUN apk add --no-cache git gcc musl-dev make curl

# Install templ tool
RUN go install github.com/a-h/templ/cmd/templ@latest

# Download and install Tailwind CSS standalone CLI
RUN curl -sLO https://github.com/tailwindlabs/tailwindcss/releases/download/v3.3.5/tailwindcss-linux-x64 \
  && chmod +x tailwindcss-linux-x64 \
  && mv tailwindcss-linux-x64 /usr/local/bin/tailwindcss

# Set working directory
WORKDIR /app

# Copy go.mod and go.sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the source code and Makefile
COPY . .

# Build CSS using tailwind
RUN make css

# Generate templ files and build the application
RUN make build

# Final stage
FROM alpine:3.19

# Install ca-certificates for HTTPS
RUN apk --no-cache add ca-certificates

WORKDIR /app

# Copy the binary from the builder stage
COPY --from=builder /app/app .

# Copy static assets and templates
COPY --from=builder /app/css ./css
COPY --from=builder /app/static ./static
COPY --from=builder /app/templates ./templates

# Expose the port the app runs on
EXPOSE 8080

# Set environment variables
ENV GO_ENV=production

# Run the application
CMD ["./app"]
