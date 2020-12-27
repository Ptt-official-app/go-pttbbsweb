#!/bin/bash

go build ./...
gotest -v ./... -run TestQueueCommentDBCS
gotest -v ./... -cover -tags noqueue
