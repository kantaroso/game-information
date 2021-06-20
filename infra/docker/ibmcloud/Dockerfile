FROM node:12.12.0-alpine

WORKDIR /usr/src/app

RUN apk update && \
    apk add curl && \
    curl -fsSL https://clis.cloud.ibm.com/install/linux | sh && \
    ibmcloud cf install

CMD ["/bin/sh"]
