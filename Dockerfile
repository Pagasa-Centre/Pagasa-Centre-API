# Build stage
FROM golang:1.24.1-alpine AS builder

WORKDIR /app

RUN apk update && apk add --no-cache git curl

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o server ./cmd/main.go

# Final stage
FROM alpine:latest

WORKDIR /root/

COPY --from=builder /app/server .
COPY --from=builder /app/internal/config ./internal/config
COPY --from=builder /app/.env .env

EXPOSE 8080

CMD ["./server"]