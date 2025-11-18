FROM golang:1.25 AS builder
WORKDIR /app

# Primero dependencias → mejor cache
COPY go.mod go.sum ./
RUN go mod download

# Ahora copiamos el código
COPY . .

RUN go build -o hotels_api main.go

# Imagen final
FROM debian:stable-slim
WORKDIR /app

# Copiamos el binario
COPY --from=builder /app/hotels_api .

EXPOSE 8082
CMD ["./hotels_api"]
