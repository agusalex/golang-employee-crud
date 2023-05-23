#!/bin/bash

if [ "$#" -ne 2 ]; then
    echo "Usage: ./kube.sh GITLAB_USERNAME GITLAB_EMAIL"
    exit 1
fi

GITLAB_USERNAME=$1
GITLAB_EMAIL=$2

read -sp 'Enter your GitLab password/token: ' GITLAB_PASSWORD

echo $GITLAB_PASSWORD | docker login registry.gitlab.com -u $GITLAB_USERNAME --password-stdin

kubectl config use-context docker-desktop

kubectl create secret docker-registry gitlab-registry \
--docker-server=registry.gitlab.com \
--docker-username=$GITLAB_USERNAME \
--docker-password=$GITLAB_PASSWORD \
--docker-email=$GITLAB_EMAIL

unset GITLAB_PASSWORD



sh deploy.sh

