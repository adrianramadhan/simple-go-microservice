# Menggunakan image golang sebagai base
FROM golang:1.22.2-alpine

# Set working directory
WORKDIR /app

# Copy go.mod dan go.sum
COPY go.mod go.sum ./

# Install dependencies
RUN go mod download

# Copy source code
COPY . .

# Build binary
RUN CGO_ENABLED=0 GOOS=linux go build -o api-gateway .

# Stage kedua untuk image yang lebih kecil
FROM alpine:latest

# Copy binary dari builder stage
COPY --from=builder /app/api-gateway .

# Jalankan aplikasi
ENTRYPOINT ["/api-gateway"]
