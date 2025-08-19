fmt:
	gofmt -w .

lint: fmt
	golangci-lint config verify
	golangci-lint run --fix

test:
	go clean -testcache
	ROOT_DIR=${PWD} go test ./...

benchmark:
	ROOT_DIR=${PWD} go test -bench=. -benchmem ./...
