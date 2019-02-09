#!/bin/bash

cd "$(dirname $0)"
DIRS=". conf format log mock"
set -e
for subdir in $DIRS; do
  pushd $subdir
  go vet
  popd
done