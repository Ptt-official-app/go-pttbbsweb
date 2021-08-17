#!/bin/bash

go build ./...
gotest -v ./... -p 1 -run TestQueueCommentDBCS -tags queue
gotest -v ./... -p 1 -cover
