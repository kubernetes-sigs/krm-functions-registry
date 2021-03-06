# render-helm-chart: Kustomize Inline Values

### Overview

This example demonstrates how to declaratively invoke the `render-helm-chart`
function with kustomize using the `valuesInline` field.

### Function invocation

To use the function with kustomize, you can specify the `functionConfig`
in your kustomization's `generators` field. This example uses inline values
to use instead of the default values accompanying the chart:

kustomization.yaml:
```yaml
generators:
  - |-
    apiVersion: v1alpha1
    kind: RenderHelmChart
    metadata:
      name: demo
      annotations:
        config.kubernetes.io/function: |
          container:
            network: true
            image: us.gcr.io/k8s-artifacts-prod/krm-functions/render-helm-chart:unstable
    helmCharts:
    - chartArgs:
        name: ocp-pipeline
        version: 0.1.16
        repo: https://bcgov.github.io/helm-charts
      templateOptions:
        releaseName: moria
        namespace: mynamespace
        values:
          valuesInline:
            releaseNamespace: ""
            rbac:
              create: true
              rules:
                - apiGroups: [""]
                  verbs: ["*"]
                  resources: ["*"]
```

Then, to build the kustomization with kustomize v4:

```shell
kustomize build --enable-alpha-plugins --network .
```

### Expected result

You should also be able to find the line `def releaseNamespace = ""` somewhere
in your output, as well as the following: 

```yaml
apiVersion: rbac.authorization.k8s.io/v1
kind: Role
metadata:
  name: moria-ocp-pipeline
  namespace: mynamespace
rules:
- apiGroups:
  - ""
  resources:
  - '*'
  verbs:
  - '*'
---
apiVersion: rbac.authorization.k8s.io/v1
kind: RoleBinding
metadata:
  name: moria-ocp-pipeline
  namespace: mynamespace
roleRef:
  apiGroup: rbac.authorization.k8s.io
  kind: Role
  name: moria-ocp-pipeline
subjects:
- kind: ServiceAccount
  name: jenkins
  namespace: mynamespace
```

which demonstrates that the correct values provided via `valuesInline` were used.