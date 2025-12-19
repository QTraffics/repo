.PHONY: fmt_install
fmt_install:
	go install -v mvdan.cc/gofumpt@latest
	go install -v github.com/daixiang0/gci@latest


.PHONY: lint_install
lint_install:
	go install -v github.com/golangci/golangci-lint/v2/cmd/golangci-lint@v2.4.0


.PHONY: fmt
fmt: script-fix
	@script/test.py fmt

.PHONY: lint
lint: fmt script-fix
	@script/test.py lint

.PHONY: test
test:
	@script/test.py test

.PHONY: test_all
test_all:
	@script/test.py fmt lint test

.PHONY: generate
generate:
	go generate ./...

.PHONY: script-fix
script-fix:
	@chmod +x ./script/*.py
