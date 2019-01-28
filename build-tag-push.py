#!/usr/bin/python

import os
import subprocess
import time
import yaml
import re
import argparse
import json

parser = argparse.ArgumentParser(description='Build images and push them, return a ECS valid file.')
parser.add_argument('-u', dest='user',
                    help='Dockerhub user.', required=True)
parser.add_argument('-i', dest='input',
                    help='Input file for docker-compose.y(a)ml', required=True)
parser.add_argument('-o', dest='output',
                    help='Output file for docker-compose.y(a)ml', required=True)

args = parser.parse_args()
project_name = os.path.basename(os.path.realpath("."))

user_name = args.user
input_file = args.input
input_dir = project_name
output_file = args.output

print "DOCKERHUB_USER=%s" % user_name

# Generate version number for built
version = str(int(time.time()))

if input_file == output_file == "docker-compose.yml":
    print "I will not clobber your docker-compose.yml file."
    print "Please unset DOCKER_COMPOSE_YML or set it to something else."
    exit(1)

print "Input file: {}".format(input_file)
print "Input dir: {}".format(input_dir)
print "Output file: {}".format(output_file)

# Execute "docker-compose build" and abort if it fails.
subprocess.check_call(["docker-compose", "-f", input_file, "build"])

# Load the services from the input docker-compose.yml file.
stack = yaml.load(open(input_file))

# Modified the script so that it handle version 2 file.
if "services" not in stack:
    print "This script expect a `services` node in docker-compose.yml"
    exit(1)

services = stack["services"]
# Iterate over all services that have a "build" definition
# Tag them, and initiate a push in the background.
push_operations = dict()
for service_name, service in services.items():
    if "build" in service:
        compose_image = "{}_{}".format(input_dir, service_name)
        hub_image = "{}/{}_{}:{}".format(user_name, input_dir, service_name, version)
        print("compose_image: {}, hub_image: {}".format(compose_image, hub_image))
        # Re-tag the image so that it can be uploaded to the the Docker Hub.
        subprocess.check_call(["docker", "tag", compose_image, hub_image])
        # Spawn "docker push" to upload the image.
        push_operations[service_name] = subprocess.Popen(["docker", "push", hub_image])
        # Replace the "build" definition by an "image" definition,
        # using the name of the image on the Docker Hub.
        del service["build"]
        service["image"] = hub_image
# Wait for push operations to complete.
for service_name, popen_object in push_operations.items():
    print "Waiting for {} push to complete...".format(service_name)
    popen_object.wait()
    print "Done."

# Write the new docker-compose.yml file.
with open(output_file, "w") as f:
    yaml.safe_dump(stack, f, default_flow_style=False)

# yaml that is produced is a bit buggy.
fh = open(output_file, "r+")
lines = map(lambda a: re.sub(r"^\s{4}-", "      -", a), fh.readlines())
fh.close()
with open(output_file, "w") as f:
    f.writelines(lines)

print "Wrote new compose file."
print "COMPOSE_FILE={}".format(output_file)