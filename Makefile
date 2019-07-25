###################### Makefile ######################
#
# Edit this file with care, as it is also being used by our CI/CD Pipeline
# For usage information check README.md
#
# Parts of this makefile are based upon github.com/kolide/kit
#

export NAME		:= argument-analysis-research
export REPO		:= canonical-debate-lab
export GIT_HOST	:= github.com
export REGISTRY	?= eu.gcr.io
export PROJECT		?= argument-analysis-research
export RBE_PROJECT		?= argument-analysis-research

export PATH 		:= $(GOPATH)/bin:$(PATH)

export NAMESPACE	?= research

-include .env
TEAMVAULT_SM ?= ~/.teamvault-sm.json

CMD ?= //...

include helpers/make_version

.PHONY: build

# test entire repo
gotest:
	@go test -cover -race $(shell go list ./... | grep -v /vendor/)

test:
	@go get github.com/onsi/ginkgo/ginkgo
	@ginkgo -r -race

gazelle:
	# bazel run //:gazelle

build: gazelle
	bazel build --workspace_status_command=./helpers/status.sh $(CMD)

docker: gazelle
	bazel run --workspace_status_command=./helpers/status.sh  $(CMD):image -- --norun

push: gazelle
	bazel run --workspace_status_command=./helpers/status.sh //:push

run: gazelle
	bazel run \
	--workspace_status_command=./helpers/status.sh $(CMD)

kube: gazelle
	bazel run \
	--verbose_failures \
	--host_force_python=PY2 \
	--workspace_status_command=./helpers/status.sh $(CMD)

# install passed in tool project
install:
	GOBIN=$(GOPATH)/bin go install cmd/$(NAME)/*.go

# run specified tool from code
dev: generate
	@go run -ldflags $(KIT_VERSION) cmd/$(NAME)/*.go \
	-debug

# format entire repo (excluding vendor)
format:
	@go get golang.org/x/tools/cmd/goimports
	@find . -type f -name '*.go' -not -path './vendor/*' -exec gofmt -w "{}" +
	@find . -type f -name '*.go' -not -path './vendor/*' -exec goimports -w "{}" +

clean:
	bazel clean

# go quality checks
check: format lint vet

# vet entire repo (excluding vendor)
vet:
	@go vet $(shell go list ./... | grep -v /vendor/)

# lint entire repo (excluding vendor)
lint:
	@go get github.com/golang/lint/golint
	@golint -min_confidence 1 $(shell go list ./... | grep -v /vendor/)

# errcheck entire repo (excluding vendor)
errcheck:
	@go get github.com/kisielk/errcheck
	@errcheck -ignore '(Close|Write)' $(shell go list ./... | grep -v /vendor/)

cover:
	@go get github.com/haya14busa/goverage
	@go get github.com/schrej/godacov
	goverage -v -coverprofile=coverage.out $(shell go list ./... | grep -v /vendor/)

generate:
	@go get github.com/maxbrunsfeld/counterfeiter
	@go generate ./...
