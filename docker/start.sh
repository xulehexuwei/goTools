#!/bin/bash

docker stop ask_baike
docker rm ask_baike

docker run -d --name ask_baike --restart=always \
-p 2345:8080 \
service-kbqa:v$1
