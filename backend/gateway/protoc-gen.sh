#!/bin/bash

set -e

cd ./pkg
services=("auth" "product" "order")
for item in "${services[@]}"; do
    cd ${item}/proto
	protoc ./*.proto --go_opt=paths=source_relative --go-grpc_opt=paths=source_relative --go_out=../pb --go-grpc_out=../pb
    cd ../..
done