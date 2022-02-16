#! /bin/bash
#
#  Copyright 2022 The Kubernetes Authors
#
#  Licensed under the Apache License, Version 2.0 (the "License");
#  you may not use this file except in compliance with the License.
#  You may obtain a copy of the License at
#
#       http://www.apache.org/licenses/LICENSE-2.0
#
#  Unless required by applicable law or agreed to in writing, software
#  distributed under the License is distributed on an "AS IS" BASIS,
#  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
#  See the License for the specific language governing permissions and
#  limitations under the License.

function build_image {
  registry="us.gcr.io/k8s-artifacts-prod/krm-functions"
  fn_name=`basename "${1}"`
  docker build -t "${registry}/${fn_name}:unstable" -f "${1}"/Dockerfile "${1}"
}

for d in */ ; do
  if [ $d != "scripts/" ]; then
    for e in "$d"*/ ; do
      fn_name=`basename "${e}"`
      if [ $fn_name = $1 ] || [ -z "${1}" ]; then
        echo building function $fn_name...
        build_image "${e::${#e}-1}"
      fi
    done
  fi
done
