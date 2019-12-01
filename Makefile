.PHONY: test

export GO111MODULE = on

test:
	go test -race -cover ./...