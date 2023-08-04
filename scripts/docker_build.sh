#!/bin/bash

branch=`git branch|grep '^*'|sed 's/^\* //g'|sed -E 's/^\(HEAD detached at //g'|sed -E 's/\)$//g'`
project=`basename \`pwd\``

GO_PTTBBS_VERSION=`grep go-pttbbs go.mod|awk '{print $2}'`
echo "GO_PTTBBS_VERSION: ${GO_PTTBBS_VERSION}"

docker build -t ${project}:${branch} -f docker/Dockerfile --build-arg GO_PTTBBS_VERSION=${GO_PTTBBS_VERSION} .
