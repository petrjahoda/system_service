#!/usr/bin/env bash
cd linux
upx system_service_linux
cd ..
#docker pull alpine:latest
docker rmi -f petrjahoda/system_service:latest
docker build -t petrjahoda/system_service:latest .
docker push petrjahoda/system_service:latest

docker rmi -f petrjahoda/system_service:2020.4.1
docker build -t petrjahoda/system_service:2020.4.1 .
docker push petrjahoda/system_service:2020.4.1
