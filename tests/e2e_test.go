// Copyright 2021 The Kubernetes Authors
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
	"testing"

	"github.com/GoogleContainerTools/kpt/pkg/test/runner"
)

// TestExamples runs all in-tree examples in the ../krm-functions/sig-cli directory as tests.
//
// It expects each subdirectory to have a '.expected' folder with the following files:
//  - 'config.yaml' has the configuration, containing the following fields:
//    - 'exitCode': The expected exit code (default 0)
//    - 'skip': If set to true, this test will be skipped (default false)
//  - 'exec.sh' is a script to run in the current directory. For example, to test a kustomization
//    file output, this can be something like `kustomize build` > resources.yaml.
//  - 'diff.patch' is the expected diff output between original directory files and
//    files after exec script running.
//
func TestExamples(t *testing.T) {
	cases, err := runner.ScanTestCases("../krm-functions")
	if err != nil {
		t.Fatalf(err.Error())
	}
	for _, c := range *cases {
		c := c
		c.Config.ImagePullPolicy = "never"
		t.Run(c.Path, func(t *testing.T) {
			r, err := runner.NewRunner(t, c, c.Config.TestType)
			if err != nil {
				t.Fatalf("error creating test runner: %s", err)
			}
			if r.Skip() {
				t.Skip()
			}
			err = r.Run()
			if err != nil {
				t.Fatalf("error running test: %s", err)
			}
		})
	}
}
