#!/bin/bash

# This script will tag the current commit with the version number
# and push the tag to the remote repository.
# usage: gtag [tagname]

# To activate this script, run ./tag.sh from the command line.
# then you can run gtag from anywhere in the repository.

alias gtag='function _gtag(){ git tag -a v$1 -m "Version $1"; git push origin --tags; };_gtag'
