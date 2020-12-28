#!/bin/bash

if [ "$1" == "" ]; then
    echo "usage: swagger.sh [host]"
    exit 255
fi

host=$1

apidoc/yamltojson.py apidoc/template.yaml apidoc/template.json
flaskswagger apidoc:app --host ${host} --base-path / --out-dir swagger --from-file-keyword=swagger_from_file --template ./apidoc/template.json

docker container stop swagger-go-openbbsmiddleware
docker container rm swagger-go-openbbsmiddleware
docker run -itd --restart always --name swagger-go-openbbsmiddleware -p 127.0.0.1:5000:8080 -e SWAGGER_JSON=/foo/swagger.json -v ${PWD}/swagger:/foo swaggerapi/swagger-ui
