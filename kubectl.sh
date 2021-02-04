#!usr/bin/env bash

kubectl create -f db-service.yaml,db-deployment.yaml,quoteapi-service.yaml,quoteapi-deployment.yaml,quotenet-networkpolicy.yaml  