#!/bin/bash

branch=`git branch|grep '^*'|sed 's/^\* //g'|sed -E 's/^\(HEAD detached at //g'|sed -E 's/\)$//g'`
project=`basename \`pwd\``

docker build -t ${project}:${branch} -f docker/Dockerfile --build-arg GO_PTTBBS_VERSION=v0.15.3 .
