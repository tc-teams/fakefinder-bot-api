FROM golang:1.12 as builder

ADD . /build/

WORKDIR /build

RUN go get github.com/go-telegram-bot-api/telegram-bot-api && \
    go build -o  main . 

FROM alpine

RUN adduser -S -D -H -h /app appuser

USER appuser

COPY --from=builder /build/main /app/

WORKDIR /app

EXPOSE 8080
