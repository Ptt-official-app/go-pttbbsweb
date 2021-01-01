#!/bin/bash

gotest ./... -coverprofile cover.out -tags noqueue

go tool cover -html=cover.out
