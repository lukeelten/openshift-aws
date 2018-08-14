
all: build

build:
	go get -d ./...
	go build ./cmd/openshift-aws

test:
	go get -t -d ./...
	go test ./...