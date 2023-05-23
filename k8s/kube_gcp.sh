#!/bin/bash

gcloud container clusters get-credentials kube-example-cluster --zone us-central1-a
docker build -t gcr.io/superb-firefly-387317/member-crud:latest ..
docker push gcr.io/superb-firefly-387317/member-crud:latest

kubectl config use-context gke_superb-firefly-387317_us-central1-a_kube-example-cluster

sh deploy.sh