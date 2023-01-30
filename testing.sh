#!/bin/sh
echo "docker compose -f docker-compose-local.yml up -d"
docker compose -f docker-compose-local.yml up -d 

echo "sleep 5s"
sleep 5

echo "docker exec test_go_api go test ./... -v"
docker exec test_go_api go test ./... -v

echo "docker compose down"
docker-compose down --remove-orphan

echo "docker volume rm godojotechtrain_db-volume"
docker volume rm godojotechtrain_db-volume
