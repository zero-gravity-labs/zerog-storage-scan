#!/bin/bash

cd ..
CURRENT_PATH="$( cd "$(dirname "$0")" >/dev/null 2>&1 ; pwd -P )"
echo "Switch dir to $CURRENT_PATH"

echo "Start to pull repo..."
git pull
echo "Success to pull repo"

echo "Start to build image..."
docker-compose build --no-cache
echo "Success to build image"

echo "Start to launch services..."
docker-compose up -d
echo "Success to launch services"