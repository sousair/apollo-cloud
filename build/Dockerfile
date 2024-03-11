FROM golang:1.22.1-alpine3.19 AS builder

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN go build -o /app/bin/httpserver ./cmd/httpserver/main.go

FROM scratch

COPY --from=builder /app/bin/httpserver /app/bin/httpserver

COPY .env .env

EXPOSE 80

CMD ["/app/bin/httpserver"]

