#!/usr/bin/env bash
cd linux
upx system_service_linux
cd ..
docker rmi -f petrjahoda/system_service:"$1"
docker build -t petrjahoda/system_service:"$1" .
docker push petrjahoda/system_service:"$1"