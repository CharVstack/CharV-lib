.PHONY: fmt lint

fmt:
	goimports -w ./

lint:
	go vet ./...
