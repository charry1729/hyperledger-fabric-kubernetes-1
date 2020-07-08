#!/bin/bash
docker build -t chary123/kube-fe:latest:dev -f Dockerfile.dev .
docker push chary123/kube123:latest:dev