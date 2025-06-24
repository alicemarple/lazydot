#!/usr/bin/env bash

cd ~/.cache/lazydot/pkg/
echo "Removed: $1"
rm -fr "$1"
