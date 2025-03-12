FROM golang:1.24-alpine AS builder

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum* ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

# Compile the application
RUN CGO_ENABLED=0 GOOS=linux go build -o /conversorgo ./cmd/app

# Final Image
FROM golang:1.24-alpine

# Install FFmpeg and other dependencies
RUN apk add --no-cache ffmpeg ca-certificates tzdata

# Create uploads directory
RUN mkdir -p /uploads

# Copy the Pre-built binary file from the previous stage
COPY --from=builder /conversorgo /usr/local/bin/conversorgo

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable using tail -f to keep the container running
CMD ["sh", "-c", "tail -f /dev/null"]
