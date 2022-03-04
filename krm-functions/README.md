# KRM Functions

This directory is the home for the source code for all in-tree functions that are maintained and
released by this repository. 

## Requirements

By donating a function's source code to this repo, you are donating the function to SIG-CLI. We cannot
accept 3rd party functions as in-tree functions.

## Files

Each publisher will need to create their own subdirectory `krm-functions/{PUBLISHER}/` 
in this directory to store their functions source code, under . For example, SIG-CLI sponsored functions are located
under `krm-functions/sig-cli/`.

Each function must be in its own subdirectory `krm-functions/{PUBLISHER}/{FUNCTION NAME}/`. This directory should 
contain: 
- An OWNERS file to approve code changes to the function.
- A README.md file to provide a user guide for the function.
- Source code and unit tests.
- A Dockerfile that describes how to build the function image.
- An `examples/` directory. An `examples/` directory. Examples that will serve both as examples for functions and as e2e tests. Each example should have its
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
