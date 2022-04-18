proto:
	./protoc-gen.sh

server:
	export ENVIRONMENT=DEVELOPMENT;\
	go run cmd/main.go

build:
	CGO_ENABLED=0 go build -o bin/main cmd/main.go

.PHONY: proto server build