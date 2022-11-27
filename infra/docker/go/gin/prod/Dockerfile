FROM golang:1.17.7-alpine3.15

COPY ./app/gin /var/www
WORKDIR /var/www
ENV IS_PROD=true

RUN go build main.go

ENTRYPOINT ["./main"]
