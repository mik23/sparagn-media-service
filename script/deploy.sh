#!/bin/bash

#This script is meant to be used to deploy the code remotely and from the circleCi pipelines
#Prerequirement: enable ssh and put public key into the server

# IP=93.38.115.75
# SSH_USERNAME=sparagn
# PROJECT_NAME=sparagn-media-service

if [ -z "${CIRCLE_BRANCH}" ]
then
  export CIRCLE_BRANCH=develop
fi

source docker/.env.${CIRCLE_BRANCH}

#zip folder to deploy including the service artifact, bash scripts and docker-compose
zip -r $PROJECT_NAME .

#send the zip to the server
scp -P 2222 -o UserKnownHostsFile=/dev/null -o StrictHostKeyChecking=no $PROJECT_NAME.zip $SSH_USERNAME@$IP:/home/sparagn

echo sparagn | ssh -o UserKnownHostsFile=/dev/null -o StrictHostKeyChecking=no -p 2222 $SSH_USERNAME@$IP bash -c "'
  export CIRCLE_BRANCH=${CIRCLE_BRANCH}
  if [ -d $PROJECT_NAME ]
  then
    cd $PROJECT_NAME
    ./script/stop.sh
    cd ..
  fi

 unzip -o $PROJECT_NAME.zip -d $PROJECT_NAME
 cd $PROJECT_NAME
 ./script/run.sh
'"
