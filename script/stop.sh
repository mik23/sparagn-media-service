#!/bin/bash

PROJECT_NAME=sparagn-media-service
rootPath=$PROJECT_NAME
cd $rootPath

echo "Stopping services... please be patient!"
cd docker/ && docker-compose down

echo "Here we go stopped!"


