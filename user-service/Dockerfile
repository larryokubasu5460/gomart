# Start from official Go image
FROM golang:1.21-alpine

# Set working directory
WORKDIR /app

# Copy go mod and download deps

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

# Build the application 
RUN go build -o main .

# Command to run
CMD ["./main"]

