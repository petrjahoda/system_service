FROM alpine:latest
RUN apk update && apk upgrade && apk add bash && apk add procps && apk add nano && apk add chromium
RUN apk add chromium
WORKDIR /bin
COPY /linux /bin
ENTRYPOINT system_service_linux
HEALTHCHECK CMD ps axo command | grep dll