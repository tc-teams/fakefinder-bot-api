FROM golang:alpine as builder

ADD . /build/

WORKDIR /build

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

RUN go get github.com/go-telegram-bot-api/telegram-bot-api && \
    go build -o  main . 

FROM alpine

RUN adduser -S -D -H -h /app appuser

USER appuser

COPY --from=builder /build/main /app/

COPY ./server/tls/ /app/tls/

WORKDIR /app

CMD ["./main","bot"]
