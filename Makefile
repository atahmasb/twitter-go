
.PHONY: test
test:
	go test -coverprofile=cover.out ./...

.PHONY: coverage
coverage: test
	go tool cover -html=cover.out

.PHONY: ci
ci: check-tidy test coverage

.PHONY: check-tidy
check-tidy: 
	go mod tidy