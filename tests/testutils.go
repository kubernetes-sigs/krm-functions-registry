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

const KrmFunctionSchema = `swagger: "2.0"
info:
  title: KRM Function Metadata
  version: v1alpha1
definitions:
  KRMFunctionDefinition:
    type: object
    description: |
      KRMFunctionDefinition is metadata that defines a KRM function
      the same way a CustomResourceDefinition defines a custom resource.
    x-kubernetes-group-version-kind:
      - group: config.kubernetes.io
        kind: KRMFunctionDefinition
        version: v1alpha1
    required:
      - apiVersion
      - kind
      - spec
    properties:
      apiVersion:
        description: apiVersion of KRMFunctionDefinition. i.e. config.kubernetes.io/v1alpha1
        type: string
        enum:
          - config.kubernetes.io/v1alpha1
      kind:
        description: kind of the KRMFunctionDefinition. It must be KRMFunctionDefinition.
        type: string
        enum:
          - KRMFunctionDefinition
      spec:
        type: object
        description: spec contains the metadata for a KRM function.
        required:
          - group
          - names
          - description
          - publisher
          - versions
        properties:
          group:
            description: group of the functionConfig
            type: string
          description:
            description: brief description of the KRM function.
            type: string
          publisher:
            description: the entity (e.g. organization) that produced and owns this KRM function.
            type: string
          names:
            description: the resource and kind names for the KRM function
            type: object
            required:
            - kind
            properties:
              kind:
                description: the Kind of the functionConfig
                type: string
          versions:
            description: the versions of the functionConfig
            type: array
            items:
              type: object
              required:
                - name
                - schema
                - idempotent
                - runtime
                - usage
                - examples
                - license
              properties:
                name:
                  description: Version of the functionConfig
                  type: string
                schema:
                  description: a URI pointing to the schema of the functionConfig
                  type: object
                  required:
                    - openAPIV3Schema
                  properties:
                    openAPIV3Schema:
                      description: openAPIV3Schema is the OpenAPI v3 schema to use for validation
                      # kube-openapi validation doesn't support references, and inlining this ref is extremely tedious
                      # $ref: "#/definitions/io.k8s.apiextensions-apiserver.pkg.apis.apiextensions.v1.JSONSchemaProps"
                idempotent:
                  description: If the function is idempotent.
                  type: boolean
                usage:
                  description: |
                    A URI pointing to a README.md that describe the details of how to
                    use the KRM function. It should at least cover what the function
                    does and what functionConfig does it support and it should give
                    detailed explanation about each field in the functionConfig.
                  type: string
                examples:
                  description: |
                    A list of URIs that point to README.md files. At least one example
                    must be provided. Each README.md should cover an example. It
                    should at least cover how to get input resources, how to run it
                    and what is the expected output.
                  type: array
                  items:
                    type: string
                license:
                  description: The license of the KRM function.
                  type: string
                  enum:
                    - Apache 2.0
                maintainers:
                  description: |
                    The maintainers for the function. It should only be used
                    when the maintainers are different from the ones in
                    spec.maintainers. When this field is specified, it
                    override spec.maintainers.
                  type: array
                  items:
                    type: string
                runtime:
                  description: |
                    The runtime information about the KRM function. At least one of
                    container and exec must be set.
                  type: object
                  properties:
                    container:
                      description: The runtime information for container-based KRM function.
                      type: object
                      required:
                        - image
                      properties:
                        image:
                          description: The image name of the KRM function.
                          type: string
                        sha256:
                          description: |
                            The digest of the image that can be verified against. It
                            is required only when the image is using semver.
                          type: string
                        requireNetwork:
                          description: If network is required to run this function.
                          type: boolean
                        requireStorageMount:
                          description: If storage mount is required to run this function.
                          type: boolean
                    exec:
                      description: The runtime information for exec-based KRM function.
                      type: object
                      required:
                        - platform
                      properties:
                        platforms:
                          description: Per platform runtime information.
                          type: array
                          items:
                            type: object
                            required:
                              - bin
                              - os
                              - arch
                              - uri
                              - sha256
                            properties:
                              bin:
                                description: The binary name.
                                type: string
                              os:
                                description: The target operating system to run the KRM function.
                                type: string
                                enum:
                                  - linux
                                  - darwin
                                  - windows
                              arch:
                                description: The target architecture to run the KRM function.
                                type: string
                                enum:
                                  - amd64
                                  - arm64
                              uri:
                                description: The location to download the binary.
                                type: string
                              sha256:
                                description: The degist of the binary that can be used to verify the binary.
                                type: string
          home:
            description: A URI pointing the home page of the KRM function.
            type: string
          maintainers:
            description: The maintainers for the function.
            type: array
            items:
              type: string
          tags:
            description: |
              The tags (or keywords) of the function. e.g. mutator, validator,
              generator, prefix, GCP.
            type: array
            items:
              type: string
paths: {}
`
