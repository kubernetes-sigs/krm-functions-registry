# Publishers

This directory is the home for all published KRM functions metadata. Each publisher should create their own 
subdirectory containing:

- An OWNERS file containing a list of owners of the publisher
- A README.md file with a description of who the publisher is
- A file for each published KRM function. The KRM function metadata file should be named after the function,
  `{FUNCTION-NAME}.yaml`. This file should be a kubernetes object of type KRMFunctionDefinition, 
  which is defined in the [Catalog KEP].
  
An example of this is SIG-CLI's [render-helm-chart](https://github.com/kubernetes-sigs/krm-functions-registry/tree/main/publishers/sig-cli/render-helm-chart.yaml) 
function metadata.
  
[Catalog KEP]: https://github.com/kubernetes/enhancements/tree/master/keps/sig-cli/2906-kustomize-function-catalog#function-metadata-schema