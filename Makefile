
.PHONY: test
test:
	go test -coverprofile=cover.out ./...

.PHONY: coverage
coverage: test
	go tool cover -html=cover.out

.PHONY: ci
ci: check-tidy test coverage

# Ensure there is no unused dependency being added by accident and all generated code is committed
.PHONY: check-tidy
check-tidy: 
	go mod tidy
	git diff --exit-code