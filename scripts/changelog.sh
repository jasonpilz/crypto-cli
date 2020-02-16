#!/usr/bin/env bash

set -eu

ORG=jasonpilz
REPO=crypto-cli

tfile=$(mktemp /tmp/$REPO-CHANGELOG-XXXXXX)
github-changelog-generator -org "$ORG" -repo "$REPO" >"$tfile"

echo "$tfile"
