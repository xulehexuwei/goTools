#!/bin/bash

rm -rf tmp && mkdir -p tmp

cp -r ../api tmp
cp -r ../config_log tmp
cp -r ../es tmp
cp -r ../qa tmp
cp -r ../utils tmp
cp -r ../main.go tmp
cp -r ../settings.ini tmp
cp -r ../go.mod tmp
cp -r ../go.sum tmp

docker build . -t service-kbqa:v$1

rm -rf tmp
