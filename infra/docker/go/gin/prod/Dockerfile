FROM golang:1.24.3-alpine3.21

COPY ./app/gin /var/www
WORKDIR /var/www
ENV IS_PROD=true

RUN go build main.go

ENTRYPOINT ["./main"]
