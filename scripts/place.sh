#!/usr/bin/env bash

cd ~/.cache/lazydot/pkg/
find . -name "*.tar.gz" -exec tar -xvf {} \;
find . -name "*.tar.gz" -exec rm {} \;
echo "Copied: $1"
cp -rf ./"$1" ~/.config/
