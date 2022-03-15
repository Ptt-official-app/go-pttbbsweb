#!/bin/bash

mkdir -p mand
protoc -I=pttbbs/daemon/mand --go_out=. --go_opt=Mman.proto=./mand --go-grpc_out=mand --go-grpc_opt=Mman.proto=./mand ./pttbbs/daemon/mand/man.proto
