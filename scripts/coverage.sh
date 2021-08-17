#!/bin/bash

gotest ./... -coverprofile cover.out
go tool cover -html=cover.out
