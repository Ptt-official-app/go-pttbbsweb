#!/bin/bash

if [ "$#" -lt "1" ]
then
    echo "usage: run-production.sh [ini-filename]"
    exit 255
fi

tags=production
ini_filename=$1

echo "to build: tags: ${tags} ini: %{ini_filename}"
go build -tags ${tags}
echo "to run pttbbs-backend"
pttbbs-backend -ini ${ini_filename}
