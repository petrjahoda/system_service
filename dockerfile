# FROM alpine:latest
# RUN apk update && apk upgrade && apk add bash && apk add procps && apk add nano
# WORKDIR /bin
# COPY /linux /bin
# ENTRYPOINT system_service_linux
# HEALTHCHECK CMD ps axo command | grep dll
FROM scratch
ADD /linux /
CMD ["/system_service_linux"]