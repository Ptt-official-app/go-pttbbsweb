#!/bin/bash

ini_filename=02-config.run.ini

echo "to build"
go build

echo "to run go-openbbsmiddleware"
./go-openbbsmiddleware -ini ${ini_filename}
