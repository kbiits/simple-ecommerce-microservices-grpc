#!/bin/bash

set -e

protoc ./pkg/proto/*.proto --go_out=. --go-grpc_out=.