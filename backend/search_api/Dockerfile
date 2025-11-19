FROM golang:1.25-alpine AS build

WORKDIR /app

COPY . .

RUN go mod download
RUN go build -o search_api main.go

FROM alpine:latest
WORKDIR /root/

COPY --from=build /app/search_api .
COPY .env .

CMD ["./search_api"]
