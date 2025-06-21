# Development stage
FROM golang:1.24

WORKDIR /app

# Install air for live reloading
RUN go install github.com/air-verse/air@latest

# Install swag for API documentation
RUN go install github.com/swaggo/swag/cmd/swag@latest

# Copy go mod files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy source code
COPY . .

# Expose port
EXPOSE 8080

# Run with air for development (live reload)
CMD ["air", "-c", ".air.toml"]
