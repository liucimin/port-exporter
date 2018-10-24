# Image URL to use all building/pushing image targets
IMG ?= port-exporter:1.0
all: port-exporter

# Build manager binary
port-exporter: generate fmt
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o bin/port_exporter github.com/port-exporter/cmd
	
# Run go fmt against code
fmt:
	go fmt ./pkg/... ./cmd/...
# Generate code
generate:
	go generate ./collector/... ./pkg/...
# Build the docker image
docker-build:
	docker build . -t ${IMG}
# Push the docker image
docker-push:
	docker push ${IMG}
