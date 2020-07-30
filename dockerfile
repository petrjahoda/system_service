# FROM alpine:latest
# RUN apk update && apk upgrade && apk add bash && apk add procps && apk add nano
# WORKDIR /bin
# COPY /linux /bin
# ENTRYPOINT system_service_linux
# HEALTHCHECK CMD ps axo command | grep dll

FROM alpine:latest as build
RUN apk add tzdata

FROM scratch as final
ADD /linux /
COPY --from=build /usr/share/zoneinfo /usr/share/zoneinfo
CMD ["/system_service_linux"]