#!/usr/bin/env bash

commitRegex="^([0-9a-f]{40}) Add branch ([0-9a-zA-Z\-]*).*$"

if [[ $1 =~ $commitRegex ]];
then 
    echo Pushing ${BASH_REMATCH[2]} as ${BASH_REMATCH[1]};
    git push -f origin ${BASH_REMATCH[1]}:refs/heads/${BASH_REMATCH[2]}
fi