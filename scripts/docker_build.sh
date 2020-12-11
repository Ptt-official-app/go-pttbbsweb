#!/bin/bash

branch=`git branch|grep '^*'|sed 's/^\* //g'|sed -E 's/^\(HEAD detached at //g'|sed -E 's/\)$//g'`
project=`basename \`pwd\``

docker build -t ${project}:${branch} .
