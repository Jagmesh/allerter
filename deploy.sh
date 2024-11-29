#!/bin/bash

while IFS= read -r line || [ -n "$line" ]; do
  if [[ "$line" == DOCKER_IMAGE_NAME=* ]]; then
    IMAGE_NAME=${line#*=}
  elif [[ "$line" == DOCKER_IMAGE_TAG=* ]]; then
    TAG=${line#*=}
  fi
done < .env

if [ -z "$IMAGE_NAME" ]; then
  echo "Error: DOCKER_IMAGE_NAME not found in .env"
  exit 1
fi
if [ -z "$TAG" ]; then
  echo "Error: DOCKER_IMAGE_TAG not found in .env"
  exit 1
fi

echo "Building Docker image: $IMAGE_NAME:$TAG"
docker image build . -t "$IMAGE_NAME:$TAG"
docker push "$IMAGE_NAME:$TAG"
