FROM golang:1.25 AS builder
WORKDIR /app
COPY . .
RUN go mod download
RUN go build -o user_api main.go

FROM debian:stable-slim
WORKDIR /app

COPY --from=builder /app/user_api .

EXPOSE 8080
CMD ["./user_api"]
