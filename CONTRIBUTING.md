# Contributing Guidelines

Welcome to Kubernetes. We are excited about the prospect of you joining our [community](https://git.k8s.io/community)! The Kubernetes community abides by the CNCF [code of conduct](code-of-conduct.md). Here is an excerpt:

_As contributors and maintainers of this project, and in the interest of fostering an open and welcoming community, we pledge to respect all people who contribute through reporting issues, posting feature requests, updating documentation, submitting pull requests or patches, and other activities._

## Getting Started

### Repo layout

```
├── publishers # Home for all functions metadata
│   ├── kustomize
│   │   ├── functions
│   │   │   ├── fn-foo.yaml
│   │   │   ├── fn-bar.yaml
│   │   │   └── README.md
│   │   ├── catalogs
│   │   │   ├── v20220225.yaml
│   │   │   ├── v20220101.yaml
│   │   │   └── README.md
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

To contribute KRM function source code, so that it can be managed and released by this repo, follow the instructions
in the [krm-functions directory]((https://github.com/kubernetes-sigs/krm-functions-registry/tree/main/krm-functions))


## Publishing KRM function metadata in-tree and out-of-tree functions

To publish KRM function metadata, independently of where the source code lives, follow the instructions 
in the [publishers directory](https://github.com/kubernetes-sigs/krm-functions-registry/tree/main/publishers).

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