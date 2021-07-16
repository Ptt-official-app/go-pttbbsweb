#!/bin/bash

go build ./...
gotest -v ./... -p 1 -run TestQueueCommentDBCS
gotest -v ./... -p 1 -cover -tags noqueue
