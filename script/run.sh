#!/bin/bash
#Script to bring up docker containers
#IT needs to be executed from the root path

DEBUG=$1

cd docker/

if [ "$DEBUG" == "--debug" ] || [ "$DEBUG" == "-d" ]
then
  optionalDebugString="-f docker-compose-debug.yml"
  echo 'debug enabled'
fi

echo 'Bringing docker up... pls wait'

docker-compose $optionalDebugString build >> /dev/null
docker-compose $optionalDebugString up -d
echo "All Deployment process done!"