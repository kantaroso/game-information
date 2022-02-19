FROM golang:1.17.7-alpine3.15

RUN apk add build-base git

WORKDIR /go/src

RUN go get -u github.com/cosmtrek/air

CMD ["air", "-c", ".air.toml"]
