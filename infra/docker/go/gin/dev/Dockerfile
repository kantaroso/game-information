FROM golang:1.24.3-alpine3.21

RUN apk add build-base git

WORKDIR /go/src

RUN go install github.com/air-verse/air@latest

CMD ["air", "-c", ".air.toml"]