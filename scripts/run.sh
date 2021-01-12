#!/bin/bash

ini_filename=02-config.run.ini

go build && ./go-openbbsmiddleware -ini ${ini_filename}
