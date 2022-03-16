#!/bin/bash
#
# Copyright 2022 The Kubernetes Authors.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#      http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.


# Usage:
#   ./release.sh $FUNCTION_NAME
#   ./release.sh $FUNCTION_NAME [patch|minor|major]
#
# If the provided function has never been released before, this script will
# release v0.1.0.
#
# Otherwise, it will do a semver bump based on the second argument (default is patch
# if no second argument is provided). This will push a tag $FUNCTION_NAME/$VERSION
# to the Github repo and will push to the release branch, e.g. $FUNCTION_NAME/v0.2.
#
# $FUNCTION_NAME should be of the form $PUBLISHER/$FUNCTION
#
# After computing the correct version number to release, the script will ask the
# user to confirm the release version and the function name before triggering
# the release.
#
# This script assumes that the remote Github repository to push the branch/tag
# is called "upstream".

set -e

REMOTE="upstream"

if [ "$#" -lt 1 ]; then
    echo "function name required as first parameter"
    exit 1
fi

FUNCTION_NAME=$1

# Check that the function source code actually exists in this repository.
for d in */ ; do
  if [ $d != "scripts/" ]; then
    for e in "$d"*/ ; do
      fn_name=${e%/}
      if [ $fn_name = $FUNCTION_NAME ]; then
        found=true
      fi
    done
  fi
done

if [ -z "$found" ]; then
    echo "error: source code for $FUNCTION_NAME not found in this repo"
    exit 1
fi

# If no release type is set, default to patch.
RELEASE_TYPE=$2
if [ -z "$RELEASE_TYPE" ]; then
    RELEASE_TYPE="patch"
fi

# Get the most recently released tag of this function.
LAST_TAG=`curl -sL http://api.github.com/repos/kubernetes-sigs/krm-functions-registry/releases | jq ".[].tag_name" | grep $FUNCTION_NAME | head -n 1`

if [ -z "$LAST_TAG" ]; then
    # function has not been released yet. We want the first release to be 0.1.0, so we set RELEASE_TYPE to minor and LAST_TAG to 0.0.0
    echo "function $FUNCTION_NAME has not been released yet, its first release will be v0.1.0"
    LAST_TAG=$FUNCTION_NAME/v0.0.0
    RELEASE_TYPE="minor"

fi

# Get the version number from the last tag.
VERSION=$(echo $LAST_TAG | sed 's:.*v::' | tr -d '"')

# Compute the next version number (following semver) and the release branch name.
if [ $RELEASE_TYPE == "patch" ]; then
  VERSION=v`echo $VERSION | awk -F. '{$3 = $3 + 1;} 1' | sed 's/ /./g'`
  RELEASE_BRANCH="$(${LAST_TAG%.*} | tr -d '"')"
  git fetch $REMOTE
  git checkout -t $REMOTE/$RELEASE_BRANCH

else
  if [ $RELEASE_TYPE == "minor" ]; then
    VERSION=v`echo $VERSION | awk -F. '{$2 = $2 + 1;} 1' | sed 's/ /./g'`

  elif [ $RELEASE_TYPE == "major" ]; then
    VERSION=v`echo $VERSION | awk -F. '{$1 = $1 + 1;} 1' | sed 's/ /./g'`

  else
    echo "error: invalid release type; must be 'patch', 'minor', or 'major'"
    exit 1
  fi

  # Create new release branch
  RELEASE_BRANCH="${LAST_TAG%/*}/${VERSION%.*}"
  git checkout -b "$RELEASE_BRANCH"

fi

# Get confirmation from user.
nl=$'\n'
read -p "Prepared to release $FUNCTION_NAME/$VERSION. Continue? [Y] or [N]${nl}" -n 1 -r
echo
if [[ $REPLY =~ ^[Yy]$ ]]; then
  echo "Triggering release..."
else
  echo "Stopped."
  exit 1
fi

# Push release branch and tag.

# assure clean workspace
echo "assuring clean workspace..."
if ! (git status | grep -q 'nothing to commit, working tree clean'); then
  echo "error: please ensure a clean workspace and run again"
  exit 1
fi

# fetch remote
echo "fetching remote..."
git fetch $REMOTE

# checkout main branch
echo "checking out main..."
git checkout main

# merge from remote main
echo "rebasing from main..."
git rebase $REMOTE/main

# assure clean workspace
echo "assuring clean workspace..."
if ! (git status | grep -q 'nothing to commit, working tree clean'); then
  echo "error: please ensure a clean workspace and run again"
  exit 1
fi

# checkout release branch
echo "checking out release branch..."
git checkout $RELEASE_BRANCH

# merge from remote main
echo "rebasing from main..."
git rebase $REMOTE/main

# push branch to remote
echo "pushing release branch to remote..."
git push -f $REMOTE $RELEASE_BRANCH

# create local release tag
echo "creating local tag..."
git tag $FUNCTION_NAME/$VERSION

# push tag to remote
echo "pushing tag to remote..."
git push $REMOTE $FUNCTION_NAME/$VERSION

# checkout main branch
echo "checking out main..."
git checkout main

# delete release branch from local
echo "deleting local release branch and tag..."
git tag --delete $FUNCTION_NAME/$VERSION
git branch -D $RELEASE_BRANCH

echo "release.sh: success."
