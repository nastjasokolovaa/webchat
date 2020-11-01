FROM golang:1.15.3-alpine3.12

RUN apk update && \
    apk add git tzdata ca-certificates

RUN git clone https://github.com/pressly/goose && \
    cd goose && \
    go mod init github.com/pressly/goose && \
    go  build -mod=vendor -tags='no_sqlite3' -v -o /bin/goose ./cmd/goose && \
    cd .. && rm -rf goose

WORKDIR /migrations
