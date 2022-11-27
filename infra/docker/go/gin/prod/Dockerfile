FROM golang:1.17.7-alpine3.15

COPY ./app/gin/main /var/www/main
COPY ./app/gin/assets/controllers /var/www/assets/controllers

WORKDIR /var/www

ENV IS_PROD=true

ENTRYPOINT ["./main"]
