#!/bin/sh
echo "docker-compose up -d go_db"
docker-compose up -d go_db

echo "go test ./... -v"
go test ./... -v

echo "docker-compose down"
docker-compose down

echo "docker volume rm godojotechtrain_db-volume"
docker volume rm godojotechtrain_db-volume