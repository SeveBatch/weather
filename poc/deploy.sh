#!/bin/bash
# This could be used for unique versioning  and tagging
# version=$(git rev-parse --short HEAD)

kubectl apply -f deployment.ym

#build new version
docker build -t weather:latest .

# tag for github container registry
docker tag weather:latest ghcr.io/sevebatch/weather:latest

# push tag
docker push ghcr.io/sevebatch/weather:latest

#update kubectl to new version
kubectl set image deployment/weather weather=weather:latest
