#!/bin/bash

COMPOSE_SERVICE=$1

OLD_CONTAINER_NAME=$(docker-compose ps "${COMPOSE_SERVICE}" | head -3 | awk '{print $1}')
OLD_CONTAINER_ID=$(docker ps -aqf "name=${OLD_CONTAINER_NAME}")

docker-compose up -d --scale "${COMPOSE_SERVICE}"=2 --no-recreate
docker kill "${OLD_CONTAINER_ID}"
docker-compose up -d --scale "${COMPOSE_SERVICE}"=1 --no-recreate
