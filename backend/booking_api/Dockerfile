FROM golang:1.25 AS builder

WORKDIR /app
COPY . .
RUN go mod download
RUN go build -o booking_api main.go

FROM debian:stable-slim
WORKDIR /app
COPY --from=builder /app/booking_api .
EXPOSE 8082
CMD ["./booking_api"]
