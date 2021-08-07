#!/bin/bash

ini_filename=docs/config/02-config.run.ini
package=github.com/Ptt-official-app/go-openbbsmiddleware/types
commit=`git rev-parse --short HEAD`
version=`git describe --tags`

go build -ldflags "-X ${package}.GIT_VERSION=${commit} -X ${package}.VERSION=${version}" && ./go-openbbsmiddleware -ini ${ini_filename}
