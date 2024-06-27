SERVICE ?= go-pttbbsweb
GOFMT ?= gofumpt -l -s
GO ?= go
BUILD_DATE ?= $(shell date -u +"%Y-%m-%dT%H:%M:%SZ")
GOFILES := $(shell find . -name "*.go")
TAGS ?=
GOPATH ?= $(shell $(GO) env GOPATH)

ifneq ($(shell uname), Darwin)
	EXTLDFLAGS = -extldflags "-static" $(null)
else
	EXTLDFLAGS =
endif

LDFLAGS ?=

all: build

.PHONY: lint
lint:
	@hash golangci-lint > /dev/null 2>&1; if [ $$? -ne 0 ]; then \
		curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.43.0; \
	fi
	golangci-lint run -v --deadline=3m

install: $(GOFILES)
	$(GO) install -v -tags '$(TAGS)' -ldflags '$(EXTLDFLAGS)-s -w $(LDFLAGS)'

build: $(SERVICE)

$(SERVICE): $(GOFILES)
	$(GO) build -v -tags '$(TAGS)' -ldflags '$(EXTLDFLAGS)-s -w $(LDFLAGS)' -o bin/$@ .

.PHONY: test
test:
	$(GO) test -p 1 -v -cover -tags $(TAGS) -coverprofile coverage.txt ./... && echo "\n==>\033[32m Ok\033[m\n" || exit 1
