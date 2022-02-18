# Publishers

This directory is the home for all published KRM functions metadata. Publishers are split into
four categories:

- communities
- companies
- github-orgs
- individuals

Each publisher should create their own subdirectory under the appropriate category. A publisher's
subdirectory should contain:

- An OWNERS file containing a list of owners of the publisher
- A directory for each published KRM function. This KRM function directory should contain a single file,
  `krm-function-metadata.yaml`. This file should be a kubernetes object of type KRMFunctionDefinition, 
  which is defined in the [Catalog KEP].
  
An example of this is SIG-CLI's [render-helm-chart](https://github.com/kubernetes-sigs/krm-functions-registry/tree/main/publishers/communities/sig-cli/render-helm-chart) 
function metadata.
  
[Catalog KEP]: https://github.com/kubernetes/enhancements/tree/master/keps/sig-cli/2906-kustomize-function-catalog#function-metadata-schema