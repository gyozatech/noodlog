.PHONY: check-coverage
check-coverage:
	go test -race -coverprofile=coverage.out -covermode=atomic