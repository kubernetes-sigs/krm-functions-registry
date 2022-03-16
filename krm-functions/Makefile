# Copyright 2022 The Kubernetes Authors.
# SPDX-License-Identifier: Apache-2.0
#
# Makefile for in-tree functions.

# Build all in-tree functions locally if function name not provided
# Otherwise, just build that function locally.
build-local:
	./scripts/build-local.sh $(FUNCTION_NAME)

# Run all unit tests for in-tree functions
unit-test:
	./scripts/unit-test.sh

# Releases the function name provided (must include publisher name)
# You can specify to release a patch, minor, or major version. The default
# is patch.
release:
	./scripts/release.sh $(FUNCTION_NAME) $(RELEASE_TYPE)
