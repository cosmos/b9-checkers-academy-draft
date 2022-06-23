#!/usr/bin/env bash

git log --pretty=oneline --no-decorate --grep="Add branch [0-9a-zA-Z\-]*" | \
    xargs -I {} ./push-branch.sh {}