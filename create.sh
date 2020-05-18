#!/usr/bin/env bash
cd linux
upx system_service_linux
cd ..
cd mac
upx system_service_mac
cd ..
cd windows
upx system_service_windows.exe
cd ..
docker pull alpine:latest
docker rmi -f petrjahoda/system_service:latest
docker build -t petrjahoda/system_service:latest .
docker push petrjahoda/system_service:latest

docker rmi -f petrjahoda/system_service:2020.2.2
docker build -t petrjahoda/system_service:2020.2.2 .
docker push petrjahoda/system_service:2020.2.2
