.PHONY: test run-dev build-run

test:
	go test -race -v -cover -covermode=atomic ./...