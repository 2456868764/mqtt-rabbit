GOOS ?= $(shell go env GOOS)

# Git information
GIT_VERSION ?= $(shell git describe --tags --always)
GIT_COMMIT_HASH ?= $(shell git rev-parse HEAD)
GIT_TREESTATE = "clean"
GIT_DIFF = $(shell git diff --quiet >/dev/null 2>&1; if [ $$? -eq 1 ]; then echo "1"; fi)
ifeq ($(GIT_DIFF), 1)
    GIT_TREESTATE = "dirty"
endif
BUILDDATE = $(shell date -u +'%Y-%m-%dT%H:%M:%SZ')

LDFLAGS = ""

# Images management
REGISTRY ?= registry.cn-hangzhou.aliyuncs.com
REGISTRY_NAMESPACE?= 2456868764
REGISTRY_USER_NAME?=""
REGISTRY_PASSWORD?=""

# Image URL to use all building/pushing image targets
BIFROMQ_ENGINE_IMG ?= ${REGISTRY}/${REGISTRY_NAMESPACE}/bifromq_engine:${GIT_VERSION}
BIFROMQ_UI_MG ?= ${REGISTRY}/${REGISTRY_NAMESPACE}/bifromq_ui:${GIT_VERSION}

## docker buildx support platform
PLATFORMS ?= linux/arm64,linux/amd64

# Get the currently used golang install path (in GOPATH/bin, unless GOBIN is set)
ifeq (,$(shell go env GOBIN))
GOBIN=$(shell go env GOPATH)/bin
else
GOBIN=$(shell go env GOBIN)
endif

LOCALBIN ?= $(shell pwd)/bin
$(LOCALBIN):
	mkdir -p $(LOCALBIN)


## Tool Binaries
SWAGGER ?= $(LOCALBIN)/swag
GOLANG_LINT ?= $(LOCALBIN)/golangci-lint
GOFUMPT  ?= $(LOCALBIN)/gofumpt


## Tool Versions
SWAGGER_VERSION ?= v1.16.1
GOLANG_LINT_VERSION ?= v1.52.2
GOFUMPT_VERSION ?= latest


# Setting SHELL to bash allows bash commands to be executed by recipes.
# This is a requirement for 'setup-envtest.sh' in the test target.
# Options are set to exit when a recipe line exits non-zero or a piped command fails.
SHELL = /usr/bin/env bash -o pipefail
.SHELLFLAGS = -ec

##@ General

# The help target prints out all targets with their descriptions organized
# beneath their categories. The categories are represented by '##@' and the
# target descriptions by '##'. The awk commands is responsible for reading the
# entire set of makefiles included in this invocation, looking for lines of the
# file as xyz: ## something, and then pretty-format the target and help. Then,
# if there's a line with ##@ something, that gets pretty-printed as a category.
# More info on the usage of ANSI control characters for terminal formatting:
# https://en.wikipedia.org/wiki/ANSI_escape_code#SGR_parameters
# More info on the awk command:
# http://linuxcommand.org/lc3_adv_awk.php

.PHONY: help
help: ## Display this help.
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n"} /^[a-zA-Z_0-9-]+:.*?##/ { printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)


.PHONY: fmt
fmt: install-gofumpt ## Run gofumpt against code.
	$(GOFUMPT) -l -w .

.PHONY: vet
vet: ## Run go vet against code.
	@find . -type f -name '*.go'| grep -v "/vendor/" | xargs gofmt -w -s

# Run mod tidy against code
.PHONY: tidy
tidy:
	@go mod tidy

.PHONY: lint
lint: install-golangci-lint  ## Run golang lint against code
	GO111MODULE=on $(GOLANG_LINT) run ./... --timeout=30m -v  --disable-all --enable=gofumpt --enable=govet --enable=staticcheck --enable=ineffassign --enable=misspell

.PHONY: test
test: fmt vet  ## Run all tests.
	go test -coverprofile coverage.out -covermode=atomic ./...

.PHONY: echoLDFLAGS
echoLDFLAGS:
	@echo $(LDFLAGS)

.PHONY: build-engine
build-engine:
	  CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build  -o cmd/docker/arm64/bifromq_engine cmd/server/main.go
	  CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build  -o cmd/docker/amd64/bifromq_engine cmd/server/main.go
.PHONY: build-ui
build-ui:
		cd "ui" && npm run build --production
	  docker build -t ${BIFROMQ_UI_MG} ./ui

.PHONY: image-buildx-engine
image-buildx-engine: build-engine  ## Build and push docker image for the dubbo client for cross-platform support
	# copy existing Dockerfile and insert --platform=${BUILDPLATFORM} into Dockerfile.cross, and preserve the original Dockerfile

	sed -e '1 s/\(^FROM\)/FROM --platform=\$$\{BUILDPLATFORM\}/; t' -e ' 1,// s//FROM --platform=\$$\{BUILDPLATFORM\}/' cmd/docker/Dockerfile > cmd/docker/Dockerfile.cross
	- docker buildx create --name project-client-builder
	docker buildx use project-client-builder
	- docker buildx build --build-arg --push --output=type=registry --platform=$(PLATFORMS) --tag ${BIFROMQ_ENGINE_IMG} -f cmd/docker/Dockerfile.cross cmd/docker
	- docker buildx rm project-client-builder
	rm cmd/docker/Dockerfile.cross && rm -f -R cmd/docker/arm64/ &&  rm -f -R cmd/docker/amd64/

.PHONY: prebuild
prebuild:  ## prebuild project
	mkdir ./external && git clone -b feat-1.20 https://github.com/2456868764/istio.git ./external/istio && git clone -b feat-security https://github.com/2456868764/dubbo-go ./external/dubbo-go && go mod tidy






