.PHONY: default build

default: build

build:
	GOOS=js GOARCH=wasm go build -o bin/json.wasm ./cmd/wasm

dev:
	go run ./cmd/server -p 9090 -s bin/

server:
	go build -o bin ./cmd/server