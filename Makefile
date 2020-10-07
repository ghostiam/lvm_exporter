GO    := go
PROMU := $(GOPATH)/bin/promu

PREFIX              ?= $(shell pwd)
BIN_DIR             ?= $(shell pwd)

GOLANGCI_LINT := $(GOPATH)/bin/golangci-lint
GOLANGCI_LINT_VERSION ?= v1.31.0

all: format lint build test

style:
	@echo ">> checking code style"
	@! gofmt -d $(shell find . -path ./vendor -prune -o -name '*.go' -print) | grep '^'

test:
	@echo ">> running tests"
	@$(GO) test -short -race ./...

format:
	@echo ">> formatting code"
	@$(GO) fmt $(pkgs)

vet:
	@echo ">> vetting code"
	@$(GO) vet $(pkgs)

build: $(PROMU)
	@echo ">> building binaries"
	$(PROMU) build --prefix $(PREFIX)

crossbuild: $(PROMU)
	@echo ">> crossbuilding binaries"
	$(PROMU) crossbuild -v

tarball: $(PROMU)
	@echo ">> building release tarball"
	@$(PROMU) tarball --prefix $(PREFIX) $(BIN_DIR)

lint: $(GOLANGCI_LINT)
	@echo ">> running golangci-lint"
	GO111MODULE=on $(GO) list -e -compiled -test=true -export=false -deps=true -find=false -tags= -- ./... > /dev/null
	GO111MODULE=on $(GOLANGCI_LINT) run --sort-results ./...

# deps
promu: $(PROMU)

$(PROMU):
	@GOOS=$(shell uname -s | tr A-Z a-z) \
		GOARCH=$(subst x86_64,amd64,$(patsubst i%86,386,$(shell uname -m))) \
		$(GO) get -u github.com/prometheus/promu

$(GOLANGCI_LINT):
	mkdir -p $(GOPATH)/bin
	curl -sfL https://raw.githubusercontent.com/golangci/golangci-lint/$(GOLANGCI_LINT_VERSION)/install.sh \
		| sed -e '/install -d/d' \
		| sh -s -- -b $(GOPATH)/bin $(GOLANGCI_LINT_VERSION)

.PHONY: all style format build crossbuild test vet tarball promu lint $(PROMU) $(GOLANGCI_LINT)
