package db

import "time"

var (
	MONGO_HOST     = "localhost"
	MONGO_PORT     = 27017
	MONGO_DBNAME   = "devptt"
	MONGO_PROTOCOL = "mongodb"

	TIMEOUT_MILLI_TS = time.Duration(500)
)
