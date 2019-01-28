#!/bin/bash

docker-compose up -d
cd ../search
go test -v ./...
cd ../rooms
go test -v ./...
cd ../queue
go test -v ./...

cd ../test
docker-compose down