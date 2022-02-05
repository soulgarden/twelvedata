fmt:
	gofmt -w .

lint: fmt
	golangci-lint run --enable-all --fix

test:
	CONFIGOR_ENV=local ROOT_DIR=${PWD} go test ./...
