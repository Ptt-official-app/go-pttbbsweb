package types

import "time"

var (
	TIMEZONE, _ = time.LoadLocation(TIME_LOCATION)
	VERSION     = ""
	GIT_VERSION = ""
)
