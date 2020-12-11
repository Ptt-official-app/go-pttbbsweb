#!/bin/bash

go build ./...
gotest -v ./... -cover
