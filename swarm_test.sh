#! /bin/bash

DRIVER=$1
MANAGER_COUNT=$2
WORKER_COUNT=$3


for i in `seq 1 $MANAGER_COUNT`
do
    docker-machine rm -y manager$i
    docker-machine create -d $DRIVER manager$i
done

eval $(docker-machine env manager1)
docker swarm init
TOKEN_MANAGER="$(docker swarm join-token -q manager)"
TOKEN_WORKER="$(docker swarm join-token -q worker)" 

echo $TOKEN_MANAGER
echo $TOKEN_WORKER
