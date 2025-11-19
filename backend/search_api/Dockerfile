FROM golang:1.25-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o search_api main.go

FROM alpine:3.19

WORKDIR /app

COPY --from=builder /app/search_api .

EXPOSE 8081

CMD ["./search_api"]
