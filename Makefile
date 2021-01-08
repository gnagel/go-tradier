.PHONY: setup proto

setup:
	brew install pre-commit || brew update pre-commit
	pre-commit install

	brew install golangci-lint || brew upgrade golangci-lint

vendor:
	go mod vendor

test:
	go test -mod=vendor -v ./...
