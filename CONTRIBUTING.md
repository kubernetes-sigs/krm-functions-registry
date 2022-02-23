# Contributing Guidelines

Welcome to Kubernetes. We are excited about the prospect of you joining our [community](https://git.k8s.io/community)! The Kubernetes community abides by the CNCF [code of conduct](code-of-conduct.md). Here is an excerpt:

_As contributors and maintainers of this project, and in the interest of fostering an open and welcoming community, we pledge to respect all people who contribute through reporting issues, posting feature requests, updating documentation, submitting pull requests or patches, and other activities._

## Getting Started

### Repo layout

```
├── publishers # Home for all functions metadata
│   ├── kustomize
│   │   ├── fn-foo.yaml
│   │   ├── fn-bar.yaml
│   │   ├── README.md
│   │   └── OWNERS # OWNERS of the publisher
│   ├── kubeflow
│   ├── sig-cli
│   └── OWNERS # OWNERS to approve new publishers
├── krm-functions # Home for in-tree functions source code
│   ├── Makefile
│   ├── kustomize
│   │   ├── fn-foo
│   │   ├── README.md
│   │   └── OWNERS # OWNERS to approve code change to the function
│   └── sig-cli
├── site 
├── Makefile
└── OWNERS
```

## Contributing in-tree KRM function source code

This section discusses how to contribute KRM function source code, so that it can be managed and released by this repo.

For each in-tree function, the implementation is in the `krm-functions/` directory. Each publisher will need to create their own subdirectory `krm-functions/{PUBLISHER}/` 
in this directory to store their functions source code. For example, SIG-CLI sponsored functions are located
under `krm-functions/sig-cli/`.

Each function must be in its own subdirectory `krm-functions/{PUBLISHER}/{FUNCTION NAME}/`. This directory should 
contain: 
- An OWNERS file to approve code changes to the function.
- A README.md file to provide a user guide for the function.
- Source code and unit tests.
- A Dockerfile that describes how to build the function image.
- An `examples/` directory. Examples that will serve both as examples for functions and as e2e tests. Each example should have its
  own subdirectory `krm-functions/{PUBLISHER}/{FUNCTION NAME}/examples/{EXAMPLE_NAME}/`, and this directory should contain:
  - A README.md file that serves as a guide for the example.
  - A subdirectory `.expected`. This should contain two files:
    - `exec.sh`: A script that will run your example. This script will be run on the example directory it is in. This can
      be something as simple as `kustomize build --enable-alpha-plugins > resources.yaml`.
    - `diff.patch`: This file should contain the expected diff between the original example directory files and the files 
      after `exec.sh` is run.
  - Any additional files needed for your examples to run. For example, if you are running `kustomize build` in your `exec.sh`
    script, you will need a kustomization file. 
        
An example of this is SIG-CLI's [render-helm-chart](https://github.com/kubernetes-sigs/krm-functions-registry/tree/main/krm-functions/sig-cli/render-helm-chart) 
function. 

## Publishing KRM function metadata in-tree and out-of-tree functions

This section describes how to publish KRM function metadata, independent of where the source code lives. 

For each function's metadata, the files are in the `publishers/` directory. This directory is the home for all published KRM functions metadata. Each publisher should create their own 
subdirectory containing:

- An OWNERS file containing a list of owners of the publisher
- A README.md file with a description of who the publisher is
- A file for each published KRM function. The KRM function metadata file should be named after the function,
  `{FUNCTION-NAME}.yaml`. This file should be a kubernetes object of type KRMFunctionDefinition, 
  which is defined in the [Catalog KEP].
    
An example of this is SIG-CLI's [render-helm-chart](https://github.com/kubernetes-sigs/krm-functions-registry/tree/main/publishers/sig-cli/render-helm-chart.yaml) 
function metadata.

## General Kubernetes Contributing docs

We have full documentation on how to get started contributing here:

- [Contributor License Agreement](https://git.k8s.io/community/CLA.md) Kubernetes projects require that you sign a Contributor License Agreement (CLA) before we can accept your pull requests
- [Kubernetes Contributor Guide](https://git.k8s.io/community/contributors/guide) - Main contributor documentation, or you can just jump directly to the [contributing section](https://git.k8s.io/community/contributors/guide#contributing)
- [Contributor Cheat Sheet](https://git.k8s.io/community/contributors/guide/contributor-cheatsheet) - Common resources for existing developers

## Mentorship

- [Mentoring Initiatives](https://git.k8s.io/community/mentoring) - We have a diverse set of mentorship programs available that are always looking for volunteers!

[Catalog KEP]: https://github.com/kubernetes/enhancements/tree/master/keps/sig-cli/2906-kustomize-function-catalog#function-metadata-schema

## Contact Information
- [Slack channel](https://kubernetes.slack.com/messages/sig-cli-krm-functions)