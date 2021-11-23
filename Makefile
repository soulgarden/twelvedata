fmt:
	gofmt -w .

lint: fmt
	golangci-lint run --enable-all --fix
