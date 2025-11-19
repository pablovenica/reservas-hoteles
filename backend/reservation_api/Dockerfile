# ---------- BUILDER ----------
FROM golang:1.25 AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o reservations_api main.go

# ---------- FINAL ----------
FROM debian:stable-slim
WORKDIR /app

# Necesario para llamadas HTTP/HTTPS entre microservicios
RUN apt-get update && apt-get install -y ca-certificates && update-ca-certificates

COPY --from=builder /app/reservations_api .

EXPOSE 8083

CMD ["./reservations_api"]

