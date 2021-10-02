FROM node:12.12.0-alpine

WORKDIR /usr/src/app

RUN apk update && \
    apk add git && \
    apk add openssh && \
    npm install -g npm && \
    npm install -g @vue/cli

EXPOSE 9050

CMD ["/bin/sh"]
