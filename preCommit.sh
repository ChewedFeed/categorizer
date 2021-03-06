#!/usr/bin/env bash
git stash -q --keep-index

go test github.com/chewedfeed/categorizer
RESULT=$?
git stash pop -q

[ $RESULT -ne 0 ] && exit 1

exit 0
