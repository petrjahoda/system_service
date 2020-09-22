[![developed_using](https://img.shields.io/badge/developed%20using-Jetbrains%20Goland-lightgrey)](https://www.jetbrains.com/go/)

![GitHub](https://img.shields.io/github/license/petrjahoda/system_service) 
![GitHub last commit](https://img.shields.io/github/last-commit/petrjahoda/system_service)
![GitHub issues](https://img.shields.io/github/issues/petrjahoda/system_service)

![GitHub language count](https://img.shields.io/github/languages/count/petrjahoda/system_service)
![GitHub top language](https://img.shields.io/github/languages/top/petrjahoda/system_service)
![GitHub repo size](https://img.shields.io/github/repo-size/petrjahoda/system_service)


![Docker Pulls](https://img.shields.io/docker/pulls/petrjahoda/system_service)
![Docker Image Size (latest by date)](https://img.shields.io/docker/image-size/petrjahoda/system_service?sort=date)

![developed_using](https://img.shields.io/badge/database-PostgreSQL-red) ![developed_using](https://img.shields.io/badge/runtime-Docker-red)




# System Service

## Description
Go service automatically creates database on its first run, check database size and disc free space. Sends email, when disc space is low.

## Installation Information
Install under docker runtime using [this dockerfile image](https://github.com/petrjahoda/system/tree/master/latest) with this command: ```docker-compose up -d```

## Implementation Information
Check the software running with this command: ```docker stats```. <br/>
System_service has to be running. No need to make a specific setup.

## Developer Information
Use software only as a [part of a system](https://github.com/petrjahoda/system) using Docker runtime.<br/>
 Do not run under linux, windows or mac on its own.
 


© 2020 Petr Jahoda
