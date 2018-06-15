
all: build

build:
	go get -d ./...
	go build ./cmd/openshift-aws
	chown $(HOST_UID) openshift-aws

test:
	go get -t -d ./...
	go test ./...