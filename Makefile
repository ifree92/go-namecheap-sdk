.PHONY: test vet

check:
	go vet ./...
	go fmt ./...

test: check
	go test -v ./...

build: check
	go build github.com/namecheap/go-namecheap-sdk

deps:
	dep ensure
