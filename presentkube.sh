#!/bin/bash

echo "[*] initiate deployment/service redis: mynet"
kubectl create deployment redis --port 6379 --image=redis
kubectl expose deployment/redis --port 6379 --name redis --type NodePort

echo "[*] initiate deployment/service mongodb"
kubectl create deployment mongo --port 27017 --image=mongo
kubectl expose deployment/mongo --port 27017 --name mongodb --type NodePort
# kubectl port-forward deployment/mongo 27017:27017

echo "[*] initiate deployment/service testapi"
kubectl create deployment testapi --port 4000 --image=unknopu/testapi
kubectl expose deployment/testapi --port 4000 --name testapi --type NodePort
# kubectl expose deployment/testapi --port 4000 --name testapi --type LoadBalancer

kubectl create deployment nginx --image=nginx
kubectl expose deployment/nginx --port 80 --name nginx --type LoadBalancer

