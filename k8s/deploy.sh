#!/bin/bash

kubectl apply -f ./mysql/secret.yaml
kubectl apply -f ./mysql/persistent_volume_claim.yaml
kubectl apply -f ./mysql/deployment.yaml
kubectl apply -f ./mysql/service.yaml

kubectl apply -f ./member-crud/deployment.yaml
kubectl apply -f ./member-crud/service.yaml

kubectl get services