#!/bin/bash

ini_filename=docs/config/02-config.run.ini
package=github.com/Ptt-official-app/go-pttbbsweb/types
commit=`git rev-parse --short HEAD`
version=`git describe --tags`

go build -ldflags "-X ${package}.GIT_VERSION=${commit} -X ${package}.VERSION=${version}" && ./go-pttbbsweb -ini ${ini_filename}
