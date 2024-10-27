# Stage 1: Build the application
FROM golang:1.22.2-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o api-gateway .

# Stage 2: Create a smaller image with the binary
FROM alpine:latest

WORKDIR /root/
COPY --from=builder /app/api-gateway .

EXPOSE 8080
ENTRYPOINT ["./api-gateway"]
