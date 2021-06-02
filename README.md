# Go Namecheap SDK

[![Go Reference](https://pkg.go.dev/badge/github.com/namecheap/go-namecheap-sdk.svg)](https://pkg.go.dev/github.com/namecheap/go-namecheap-sdk)
[![Go Report Card](https://goreportcard.com/badge/github.com/namecheap/go-namecheap-sdk)](https://goreportcard.com/report/github.com/namecheap/go-namecheap-sdk)

### Getting

```
$ go get github.com/namecheap/go-namecheap-sdk
```

### Usage

Generally callers would create a `namecheap.Client` and make calls off of that.

```go
import (
    "github.com/namecheap/go-namecheap-sdk"
)

// Reads environment variables
client, err := namecheap.New()

// Directly build client
client, err := namecheap.NewClient(username, apiuser string, token string, ip string, useSandbox)
```

Calling `namecheap.New()` reads the following environment variables:

- `NAMECHEAP_USERNAME`: Username: e.g. john
- `NAMECHEAP_API_USER`: ApiUser: e.g. john
- `NAMECHEAP_TOKEN`: From https://ap.www.namecheap.com/Profile/Tools/ApiAccess
- `NAMECHEAP_IP`: Your IP (must be whitelisted)
- `NAMECHEAP_USE_SANDBOX`: Use sandbox environment

### Contributing

We appreciate feedback, issues and Pull Requests. You can build the project with `make build` in the root and run tests with `make test`.

If you're looking to run tests yourself you can configure the environmental variables and override the test records in `client_test.go`. (To make live API calls) Otherwise only mockable tests will run.

The following are contributor oriented environmental variables:

- `DEBUG`: Log all responses
- `MOCKED`: Force disable `testClient`
