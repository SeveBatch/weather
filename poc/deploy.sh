#!/bin/bash
version=$(git rev-parse --short HEAD)

#build new version
docker build -t weather:"${version}" .

# tag for github container registry
docker tag  github.com/sevebatch/weather:"${version}" ghcr.io/sevebatch/weather:"${version}"

# push tag
docker push ghcr.io/sevebatch/weather:"${version}"

#update kubectl to new version
kubectl set image deployment/weather weather=weather:"${version}"
