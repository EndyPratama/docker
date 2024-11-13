FROM golang:1.21-alpine AS builder

WORKDIR /app-test

# Copy the Go module files and download dependencies
COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o ./build/app .

CMD ["./build/app"]