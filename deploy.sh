#!/bin/bash


# This file will deploy the karaoke-app.
# The deployment is made to amazon ECS.


echo "Removing old source files."
ssh ubuntu@ec2-52-31-108-82.eu-west-1.compute.amazonaws.com "sudo rm -rf karaoke-app/backend;sudo rm karaoke-app/docker-compose.yml karaoke-app/traefik.toml; exit"
echo "Copy new source files."
scp -r ./dist/* ubuntu@ec2-52-31-108-82.eu-west-1.compute.amazonaws.com:/home/ubuntu/karaoke-app
echo "Taking down current service and re-enabling"
ssh ubuntu@ec2-52-31-108-82.eu-west-1.compute.amazonaws.com "cd karaoke-app; sudo chmod 600 acme.json; sudo docker-compose down; sudo docker-compose up -d; exit"
echo "Deployed."
