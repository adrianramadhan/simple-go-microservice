FROM golang:1.22.2-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o user-service .

FROM alpine:latest

WORKDIR /root/
COPY --from=builder /app/user-service .

EXPOSE 8080
ENTRYPOINT ["./user-service"]
