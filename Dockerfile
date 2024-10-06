# First stage: Build the Go binary
FROM golang:1.23.1-alpine AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and regenerate go.sum
COPY go.mod ./
RUN go mod tidy

# Copy the remaining source files
COPY . .

# Build the Go binary with optimizations
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o /app/blu-installment .

# Second stage: Run the Go binary in a minimal image
FROM alpine:3.18.3

# Expose port 8000 for the HTTP server
EXPOSE 8000

# Copy the binary from the builder stage
COPY --from=builder /app/blu-installment /usr/local/bin/blu-installment

# Make sure the binary is executable
RUN chmod +x /usr/local/bin/blu-installment

# Set an entrypoint or CMD to run the executable
CMD ["blu-installment"]
