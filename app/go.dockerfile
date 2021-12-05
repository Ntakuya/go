FROM golang:1.17

RUN apt update && apt upgrade -y
RUN go get -u github.com/cespare/reflex
