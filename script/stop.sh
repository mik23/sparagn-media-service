#!/bin/bash

CIRCLE_PROJECT_REPONAME=sparagn-media-service
arrIN=(${PWD//$CIRCLE_PROJECT_REPONAME/ })
rootPath=$arrIN$CIRCLE_PROJECT_REPONAME
cd $rootPath

echo "Stopping services... please be patient!"
cd docker/ && docker-compose down

echo "Here we go stopped!"


