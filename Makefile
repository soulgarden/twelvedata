fmt:
	gofmt -w .

lint: fmt
	golangci-lint run --fix

test:
	go clean -testcache
	CONFIGOR_ENV=local ROOT_DIR=${PWD} go test ./...
