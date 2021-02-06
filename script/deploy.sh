#!/bin/bash

#This script is meant to be used to deploy the code remotely and from the circleCi pipelines
#Prerequirement: enable ssh and put public key into the server

# IP=93.38.115.75
# SSH_USERNAME=sparagn
# CIRCLE_PROJECT_REPONAME=sparagn-media-service

if [[ -z $CIRCLE_BRANCH ]] || [[  $CIRCLE_BRANCH != "develop" &&  $CIRCLE_BRANCH != "master" ]];
then
  export CIRCLE_BRANCH=develop
fi

source docker/.env.${CIRCLE_BRANCH}

#zip folder to deploy including the service artifact, bash scripts and docker-compose
zip -r $CIRCLE_PROJECT_REPONAME .

#send the zip to the server
scp -P $SSH_PORT -o UserKnownHostsFile=/dev/null -o StrictHostKeyChecking=no $CIRCLE_PROJECT_REPONAME.zip $SSH_USERNAME@$IP:/home/sparagn

echo sparagn | ssh -o UserKnownHostsFile=/dev/null -o StrictHostKeyChecking=no -p $SSH_PORT $SSH_USERNAME@$IP bash -c "'
  export CIRCLE_BRANCH=${CIRCLE_BRANCH}
  if [ -d $CIRCLE_PROJECT_REPONAME ]
  then
    cd $CIRCLE_PROJECT_REPONAME
    ./script/stop.sh
    cd ..
  fi

 unzip -o $CIRCLE_PROJECT_REPONAME.zip -d $CIRCLE_PROJECT_REPONAME
 cd $CIRCLE_PROJECT_REPONAME
 ./script/run.sh
'"
