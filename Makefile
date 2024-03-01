
# Parameters
GOLANGCILINT := $(GOPATH)/bin/golangci-lint
GOLANGCILINTVERSION := 1.54.2
$(GOLANGCILINT):
	curl -fsSL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(GOPATH)/bin v$(GOLANGCILINTVERSION)


# Commands
.PHONY: test
test:
	go test ./...

.PHONY: lint
lint: $(GOLANGCILINT)
	$(GOLANGCILINT) run --timeout=3m --max-issues-per-linter=0 --max-same-issues=0 -v

.PHONY: lint-all
lint-all: $(GOLANGCILINT)
	$(GOLANGCILINT) run --timeout=3m --max-issues-per-linter=0 --max-same-issues=0
