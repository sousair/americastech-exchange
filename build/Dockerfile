FROM golang:1.21.3-alpine3.17 AS compiler

WORKDIR /app

COPY . .

ENV CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

RUN go build -o /app/bin/httpserver ./cmd/httpserver/main.go

#

FROM alpine

COPY --from=compiler /app/bin /app/bin

COPY .env .env

EXPOSE 80

ENTRYPOINT ["/app/bin/httpserver"]
