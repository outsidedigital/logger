.DEFAULT_GOAL = help

GIT_SHA := $(shell git rev-parse --short HEAD 2>/dev/null)
GIT_TAG := $(shell git tag --points-at HEAD 2>/dev/null)

REPORTS_DIR := $(if $(REPORTS_DIR),$(REPORTS_DIR:/=),target/reports)

VERSION ?= $(if $(GIT_TAG),$(GIT_TAG:v%=%),unknown$(if $(GIT_SHA), on $(GIT_SHA),))

## clean: Remove created resources.
.PHONY: clean
clean:
	rm -rf $(REPORTS_DIR)

## help: Display available targets.
.PHONY: help
help: $(MAKEFILE_LIST)
	@echo "Usage: make [target]"
	@echo
	@echo "Targets:"
	@sed -En 's/^## *([^:]+): *(.*)$$/\1\t\2/p' $< | expand -t 18

## lint: Run static analysis checks.
.PHONY: lint
lint:
	golangci-lint run

## test: Run tests and generate quality reports.
.PHONY: test
test tests: $(REPORTS_DIR)
	gotestsum --junitfile $(REPORTS_DIR)/tests.xml -f standard-quiet -- \
	-coverpkg ./... -covermode atomic -coverprofile $(REPORTS_DIR)/cover.out \
	./...
	coverage -i "$(wildcard cmd/*/*.go)" $(REPORTS_DIR)/cover.out

$(REPORTS_DIR):
	mkdir -p "$@"

## version: Display current version.
.PHONY: version
version:
	@echo 'version $(VERSION)'
