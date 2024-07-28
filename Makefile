GOLANGCI_LINT_VERSION = 1.59.1
ACTIONLINT_VERSION = 1.7.1

test:
	go test --timeout 10m -race ./...

coverage:
	go test -race -v -coverpkg=./... -coverprofile=profile.out ./...
	go tool cover -func profile.out

lint:
	go run github.com/golangci/golangci-lint/cmd/golangci-lint@v$(GOLANGCI_LINT_VERSION) run

actionlint:
	go run github.com/rhysd/actionlint/cmd/actionlint@v$(ACTIONLINT_VERSION)