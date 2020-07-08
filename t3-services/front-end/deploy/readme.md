Dockerfile 
Dockerfile.dev
docker-compose.yaml

docker-compose up =--local

docker build -t chary123/kube-fe:latest-dev -f Dockerfile.dev .
docker run container name

docker push chary123/kube-fe:latest-dev


docker run chary123/kube-fe:latest-dev