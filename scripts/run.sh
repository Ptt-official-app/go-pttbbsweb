#!/bin/bash

ini_filename=00-config.template.ini

echo "to build"
go build

echo "to run go-openbbsmiddleware"
./go-openbbsmiddleware -ini ${ini_filename}
