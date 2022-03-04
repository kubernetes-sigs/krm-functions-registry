// Copyright 2022 The Kubernetes Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package tests

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"k8s.io/kube-openapi/pkg/validation/spec"
	"k8s.io/kube-openapi/pkg/validation/strfmt"
	"k8s.io/kube-openapi/pkg/validation/validate"
	"sigs.k8s.io/kustomize/kyaml/fn/framework"
	k8syaml "sigs.k8s.io/yaml"
)

// TestValidateMetadata validates that the function metadata in each function metadata
// file of the publishers directory is a valid KRMFunctionDefinition.
func TestValidateMetadata(t *testing.T) {
	assert.NoError(t, os.Chdir("../publishers"))
	publishers, err := os.ReadDir(".")
	assert.NoError(t, err)

	for _, publisher := range publishers {
		if publisher.IsDir() {
			functionDir := filepath.Join(publisher.Name(), "functions")
			functions, err := os.ReadDir(functionDir)
			assert.NoError(t, err)

			for _, function := range functions {
				validateFunctionMetadata(t, filepath.Join(functionDir, function.Name()))
			}
		}
	}
}

func validateFunctionMetadata(t *testing.T, path string) {
	functionMetadata, err := os.ReadFile(path)
	require.NoError(t, err)

	schemaJSON, err := k8syaml.YAMLToJSON([]byte(KrmFunctionSchema))
	require.NoError(t, err)

	swagger := &spec.Swagger{}
	require.NoError(t, swagger.UnmarshalJSON(schemaJSON))

	var input map[string]interface{}
	inputJSON, err := k8syaml.YAMLToJSON(functionMetadata)
	require.NoError(t, err)
	require.NoError(t, json.Unmarshal(inputJSON, &input))

	schema := swagger.Definitions["KRMFunctionDefinition"]
	schemaValidationError := validate.AgainstSchema(&schema,
		input, strfmt.Default)

	require.NoError(t, schemaValidationError)

	var def framework.KRMFunctionDefinition
	err = k8syaml.Unmarshal(functionMetadata, &def)
	require.NoError(t, err)
}
