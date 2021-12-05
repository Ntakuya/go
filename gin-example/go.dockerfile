FROM golang:1.17

WORKDIR /go/src
COPY . .

RUN apt update && apt upgrade -y
RUN go get -u github.com/cespare/reflex
