package schema

import "time"

var (
	MONGO_HOST     = "localhost"
	MONGO_PORT     = 27017
	MONGO_DBNAME   = "devptt"
	MONGO_PROTOCOL = "mongodb"

	MAX_COMMENT_BLOCK         = 20
	MAX_COMMENT_SUMMARY_BLOCK = 50

	REDIS_HOST             = "localhost:6379"
	REDIS_TIMEOUT_MILLI_TS = time.Duration(10)
)
