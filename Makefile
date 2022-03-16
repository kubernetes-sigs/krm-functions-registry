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
all: install-tools license validate-metadata test-in-tree

## Run all unit tests and e2e tests for in-tree functions
test-in-tree: e2e-test unit-test lint

# Run all the e2e tests for in-tree functions
e2e-test: build-local $(MYGOBIN)/kustomize
	cd tests; \
    go test -v -run TestExamples
    
# Validate the metadata for all published functions
validate-metadata:
	cd tests; \
    go test -v -run TestValidateMetadata

# Run all unit tests for in-tree functions
unit-test:
	cd krm-functions && $(MAKE) unit-test

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

# Release all in-tree functions
release:
	cd krm-functions && $(MAKE) release

# Install tools needed to run tests
install-tools: \
	install-kustomize \
	install-addlicense \
	install-golangci-lint \
	install-helm

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

# Install helm v3 (needed for sig-cli/render-helm-chart)
install-helm:
ifeq (, $(shell which helm))
	curl https://raw.githubusercontent.com/helm/helm/main/scripts/get-helm-3 | bash
endif
