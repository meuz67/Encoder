# syntax=docker/dockerfile:1

FROM golang:1.26-alpine AS builder
WORKDIR /workspace

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o /app/cipher ./cmd/app

FROM alpine:latest
RUN apk add --no-cache ca-certificates
COPY --from=builder /app/cipher /usr/local/bin/cipher
WORKDIR /app

EXPOSE 8080
CMD ["/usr/local/bin/cipher"]
