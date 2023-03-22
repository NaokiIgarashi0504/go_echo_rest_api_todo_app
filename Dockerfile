FROM golang:1.19.1-alpine

WORKDIR /go/src
COPY . .

RUN go get github.com/mattn/go-sqlite3