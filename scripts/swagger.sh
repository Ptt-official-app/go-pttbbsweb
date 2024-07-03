#!/bin/bash

if [ "$1" == "" ]; then
    echo "usage: swagger.sh [host]"
    exit 255
fi

host=$1
currentDir=`pwd`

apidoc/yamltojson.py apidoc/template.yaml apidoc/template.json
flaskswagger apidoc:app --host ${host} --base-path / --out-dir swagger --from-file-keyword=swagger_from_file --template ./apidoc/template.json

docker run --rm -v ${currentDir}/swagger:/data swaggerapi/swagger-codegen-cli-v3 generate -l openapi-yaml -i file:///data/swagger.json -o /data/v3
apidoc/addsecurity.py swagger/v3/openapi.yaml
apidoc/parameters.py swagger/swagger.json swagger/v3/openapi.security.json

# docker container stop swagger-go-pttbbsweb
# docker container rm swagger-go-pttbbsweb
# docker run -itd --restart always --name swagger-go-pttbbsweb -p 127.0.0.1:5000:8080 -e SWAGGER_JSON=/foo/swagger.json -v ${PWD}/swagger:/foo swaggerapi/swagger-ui

docker container stop swagger-go-pttbbsweb-v3
docker container rm swagger-go-pttbbsweb-v3
docker run -itd --restart always --name swagger-go-pttbbsweb-v3 -p 127.0.0.1:5000:8080 -e SWAGGER_JSON=/foo/v3/openapi.security.params.json -v ${PWD}/swagger:/foo swaggerapi/swagger-ui
