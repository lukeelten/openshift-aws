all:
	go get -d ./...
	go build ./cmd/openshift-aws
	chown $(HOST_UID) openshift-aws