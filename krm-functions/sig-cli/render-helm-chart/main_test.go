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

package main

import (
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
	"sigs.k8s.io/krm-functions-registry/krm-functions/sig-cli/render-helm-chart/third_party/sigs.k8s.io/kustomize/api/builtins"
	"sigs.k8s.io/krm-functions-registry/krm-functions/sig-cli/render-helm-chart/third_party/sigs.k8s.io/kustomize/api/types"
	"sigs.k8s.io/kustomize/kyaml/yaml"
)

func TestHelmChartInflatorFunction_Config(t *testing.T) {
	tests := map[string]struct {
		functionConfig string
		expectedErr    string
		expected       []builtins.HelmChartInflationGeneratorPlugin
	}{
		"invalid function config": {
			functionConfig: `invalid`,
			expectedErr:    "missing Resource metadata",
		},
		"function config as ConfigMap": {
			functionConfig: `
apiVersion: v1
kind: ConfigMap
metadata:
  name: myMap
data:
  name: minecraft
  repo: https://itzg.github.io/minecraft-server-charts
  version: 3.1.3
  releaseName: test
`,
			expected: []builtins.HelmChartInflationGeneratorPlugin{
				{
					HelmChart: types.HelmChart{
						ChartArgs: types.ChartArgs{
							Name:    "minecraft",
							Repo:    "https://itzg.github.io/minecraft-server-charts",
							Version: "3.1.3",
						},
						TemplateOptions: types.TemplateOptions{
							ReleaseName: "test",
							Values: types.Values{
								ValuesFiles: []string{"tmp/charts/minecraft/values.yaml"},
								ValuesMerge: "override",
							},
						},
					},
				},
			},
		},
		"function config as RenderHelmChart": {
			functionConfig: `
apiVersion: v1
kind: RenderHelmChart
metadata:
  name: myRenderHelmChart
helmCharts:
- chartArgs:
    name: minecraft
    repo: https://itzg.github.io/minecraft-server-charts
    version: 3.1.3
  templateOptions:
    releaseName: test-1
- chartArgs:
    name: minecraft
    repo: https://itzg.github.io/minecraft-server-charts
    version: 3.1.3
  templateOptions:
    releaseName: test-2
`,
			expected: []builtins.HelmChartInflationGeneratorPlugin{
				{
					HelmChart: types.HelmChart{
						ChartArgs: types.ChartArgs{
							Name:    "minecraft",
							Repo:    "https://itzg.github.io/minecraft-server-charts",
							Version: "3.1.3",
						},
						TemplateOptions: types.TemplateOptions{
							ReleaseName: "test-1",
							Values: types.Values{
								ValuesFiles: []string{"tmp/charts/minecraft/values.yaml"},
								ValuesMerge: "override",
							},
						},
					},
				},
				{
					HelmChart: types.HelmChart{
						ChartArgs: types.ChartArgs{
							Name:    "minecraft",
							Repo:    "https://itzg.github.io/minecraft-server-charts",
							Version: "3.1.3",
						},
						TemplateOptions: types.TemplateOptions{
							ReleaseName: "test-2",
							Values: types.Values{
								ValuesFiles: []string{"tmp/charts/minecraft/values.yaml"},
								ValuesMerge: "override",
							},
						},
					},
				},
			},
		},
	}

	for _, tc := range tests {
		var fn helmChartInflatorFunction
		node := yaml.MustParse(tc.functionConfig)
		err := fn.Config(node)
		if tc.expectedErr == "" {
			if !assert.NoError(t, err) {
				t.FailNow()
			}
			assert.Equal(t, len(tc.expected), len(fn.plugins))
			for i := range tc.expected {
				assert.Equal(t, tc.expected[i].HelmChart, fn.plugins[i].HelmChart)
			}
		} else {
			if !assert.Error(t, err) {
				t.FailNow()
			}
			assert.Contains(t, err.Error(), tc.expectedErr)
		}
	}
}

func TestHelmChartInflatorFunction_Run(t *testing.T) {
	tmpDir, err := ioutil.TempDir("", "")
	assert.NoError(t, err)

	tests := map[string]struct {
		functionConfig string
		expected       string
	}{
		"simple": {
			functionConfig: `
apiVersion: v1
kind: RenderHelmChart
metadata:
  name: myRenderHelmChart
helmGlobals:
  chartHome: ` + tmpDir + `
helmCharts:
- chartArgs:
    name: minecraft
    repo: https://itzg.github.io/minecraft-server-charts
    version: 3.1.3
  templateOptions:
    releaseName: test-1
- chartArgs:
    name: minecraft
    repo: https://itzg.github.io/minecraft-server-charts
    version: 3.1.3
  templateOptions:
    releaseName: test-2
`,
			expected: `# Source: minecraft/templates/secrets.yaml
apiVersion: v1
kind: Secret
metadata:
  name: test-1-minecraft
  labels:
    app: test-1-minecraft
    chart: "minecraft-3.1.3"
    release: "test-1"
    heritage: "Helm"
type: Opaque
data:
  rcon-password: "Q0hBTkdFTUUh"
# Source: minecraft/templates/minecraft-svc.yaml
apiVersion: v1
kind: Service
metadata:
  name: test-1-minecraft
  labels:
    app: test-1-minecraft
    chart: "minecraft-3.1.3"
    release: "test-1"
    heritage: "Helm"
  annotations: {}
spec:
  type: ClusterIP
  ports:
  - name: minecraft
    port: 25565
    targetPort: minecraft
    protocol: TCP
  selector:
    app: test-1-minecraft
# Source: minecraft/templates/secrets.yaml
apiVersion: v1
kind: Secret
metadata:
  name: test-2-minecraft
  labels:
    app: test-2-minecraft
    chart: "minecraft-3.1.3"
    release: "test-2"
    heritage: "Helm"
type: Opaque
data:
  rcon-password: "Q0hBTkdFTUUh"
# Source: minecraft/templates/minecraft-svc.yaml
apiVersion: v1
kind: Service
metadata:
  name: test-2-minecraft
  labels:
    app: test-2-minecraft
    chart: "minecraft-3.1.3"
    release: "test-2"
    heritage: "Helm"
  annotations: {}
spec:
  type: ClusterIP
  ports:
  - name: minecraft
    port: 25565
    targetPort: minecraft
    protocol: TCP
  selector:
    app: test-2-minecraft
`,
		},
		"include CRDs": {
			functionConfig: `
apiVersion: v1
kind: RenderHelmChart
metadata:
  name: myRenderHelmChart
helmGlobals:
  chartHome: ` + tmpDir + `
helmCharts:
- chartArgs:
    name: terraform
    repo: https://helm.releases.hashicorp.com
    version: 1.0.0
  templateOptions:
    releaseName: terraforming-mars
    includeCRDs: true
`,
			expected: `kind: CustomResourceDefinition`,
		},
	}

	for _, tc := range tests {
		var fn helmChartInflatorFunction
		node := yaml.MustParse(tc.functionConfig)

		err = fn.Config(node)
		assert.NoError(t, err)

		items, err := fn.Run(nil)
		if !assert.NoError(t, err) {
			t.FailNow()
		}

		actual := ""
		for i := range items {
			actual = actual + items[i].MustString()
		}
		assert.Contains(t, actual, tc.expected)
	}
}
