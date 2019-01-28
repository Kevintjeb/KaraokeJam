#!/bin/bash


# This file will build the karaoke-app.
# First it will build the backend, then the frontend.
# This does *NOT* start the application!

python build-tag-push.py -u kevintjeb -i docker-compose.yml -o ./dist/docker-compose.yml