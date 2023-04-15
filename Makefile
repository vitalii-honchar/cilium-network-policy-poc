
.PHONY: test

test:
	go test -v ./...

gofumpt:
	go gofumpt -w .

.PHONY: build
build:
	mkdir -p build
	go build -o build -v ./... 

dockerbuild: build
	docker build -t resilience-test/worker:v1.0.0 -f Dockerfile.worker .
	docker build -t resilience-test/server:v1.1.0 -f Dockerfile.server .