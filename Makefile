default: check test build

check:
	go vet ./...
	go fmt ./...

test:
	go test -v ./...

build:
	go build github.com/namecheap/go-namecheap-sdk
