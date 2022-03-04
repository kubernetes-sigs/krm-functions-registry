# Publishers

This directory is the home for all published KRM functions metadata. Each publisher should create their own 
subdirectory containing:

- An OWNERS file containing a list of owners of the publisher
- A README.md file with a description of who the publisher is
- A `functions` directory, containing a file for each published KRM function. The KRM function metadata file should be named after the function,
  `{FUNCTION-NAME}.yaml`. This file should be a kubernetes object of type KRMFunctionDefinition, 
  which is defined in the [Catalog KEP].
- A `catalogs` directory, containing a file for each published functions Catalog. Each catalog file should be named with the date it was
  published, in the form `v{YYYY}{MM}{DD}.yaml`, e.g. `v20220225.yaml`. 
  
See [SIG-CLI functions] as an example.

[SIG-CLI functions]: (https://github.com/kubernetes-sigs/krm-functions-registry/tree/main/publishers/sig-cli)
  
[Catalog KEP]: https://github.com/kubernetes/enhancements/tree/master/keps/sig-cli/2906-kustomize-function-catalog#function-metadata-schema