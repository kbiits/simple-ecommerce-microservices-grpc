#!/bin/bash

set -e

# start all services and gateway
services=("gateway" "auth" "order" "product")
for item in "${services[@]}"; do
    cd backend/${item}
    echo "Build ${item}"
    make build
    cd ../..
done

echo "Run docker compose"
docker-compose up -d
