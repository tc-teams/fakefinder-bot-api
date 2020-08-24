FROM golang:alpine as builder

ADD . /build/

WORKDIR /build

RUN go get github.com/go-telegram-bot-api/telegram-bot-api && \
    go build -o  main . 

FROM alpine

RUN adduser -S -D -H -h /app appuser

USER appuser

ENV TELEGRAM_BOT_KEY=$TELEGRAM_BOT_KEY

COPY --from=builder /build/main /app/

WORKDIR /app

CMD ["./main","bot"]
