# Copyright 2022 The Kubernetes Authors.
# SPDX-License-Identifier: Apache-2.0
#
# Makefile for KRM functions registry.

MYGOBIN = $(shell go env GOBIN)
ifeq ($(MYGOBIN),)
MYGOBIN = $(shell go env GOPATH)/bin
endif
export PATH := $(MYGOBIN):$(PATH)

# Run all tests in this repo for in-tree and out-of-tree functions
all: install-tools license verify-metadata test-in-tree

## Run all unit tests and e2e tests for in-tree functions
test-in-tree: e2e-test unit-test lint

# Run all the e2e tests for in-tree functions
e2e-test: build-local $(MYGOBIN)/kustomize
	cd tests; \
    go test -v ./...

# Build all in-tree functions locally
build-local:
	cd krm-functions && $(MAKE) build-local

# Add project licenses to all code files in here
.PHONY: license
license: $(MYGOBIN)/addlicense
	( find . -type f -exec bash -c "$(MYGOBIN)/addlicense -y 2022 -c 'The Kubernetes Authors.' -f LICENSE_TEMPLATE {}" ";" )

# Lint all in-tree functions
lint:
	cd tests; golangci-lint run ./...
	cd krm-functions && find . -type f -name go.mod -execdir golangci-lint run ./... \;

# TODO
# Run all unit tests for in-tree functions
unit-test:

# TODO
# Verify the metadata for all in-tree and out-of-tree functions
verify-metadata:

# Install tools needed to run tests
install-tools: \
	install-kustomize \
	install-addlicense \
	install-golangci-lint

# Install kustomize
install-kustomize:
ifeq (, $(shell which kustomize))
	curl -s "https://raw.githubusercontent.com/kubernetes-sigs/kustomize/master/hack/install_kustomize.sh"  | bash; \
	mv ./kustomize $(MYGOBIN)/kustomize
endif

# Install the addlicense tool
install-addlicense:
	go install github.com/google/addlicense@v1.0.0

# Install the lint tool
install-golangci-lint:
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.44.0
