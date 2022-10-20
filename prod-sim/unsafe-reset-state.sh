#!/usr/bin/env bash

here=$(cd $(dirname "$0"); pwd -P)

echo -e node-carol'\n'sentry-alice'\n'sentry-bob'\n'val-alice'\n'val-bob \
    | xargs -I {} -P 5 \
    docker run --rm -i \
    -v $here/{}:/root/.checkers \
    checkersd_i \
    tendermint unsafe-reset-all \
    --home /root/.checkers

git checkout -- $here/kms-alice/state/checkers-1-consensus.json
git checkout -- $here/val-bob/data/priv_validator_state.json

rm $here/node-carol/config/write-file*
rm $here/sentry-alice/config/write-file*
rm $here/sentry-bob/config/write-file*
rm $here/val-alice/config/write-file*
rm $here/val-bob/config/write-file*