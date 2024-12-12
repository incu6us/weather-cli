.PHONY install:
install:
	@echo "Installing binaries..."
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	go install github.com/gotesttools/gotestfmt/v2/cmd/gotestfmt@latest
	go install go.uber.org/mock/mockgen@latest
	@echo "Done."

.PHONY test:
test:
	@echo "Testing..."
	@go test -race -json -v -coverprofile=coverage.txt ./... 2>&1 | tee /tmp/gotest.log | gotestfmt
	@echo "Done."

.PHONY lint:
lint:
	@echo "Linting..."
	@echo "golangci-lint path: $$(which golangci-lint)"
	@echo "Version: $$(golangci-lint --version)"
	@golangci-lint run
	@echo "Done."
