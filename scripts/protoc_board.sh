#!/bin/bash

mkdir -p boardd
protoc -I=pttbbs/daemon/boardd --go_out=. --go_opt=Mboard.proto=./boardd --go-grpc_out=boardd --go-grpc_opt=Mboard.proto=./boardd ./pttbbs/daemon/boardd/board.proto
