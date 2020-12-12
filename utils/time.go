package utils

import "time"

const (
	TS_TO_NANO_TS       = 1000000000                     //10^9
	TS_TO_MILLI_TS      = 1000                           //10^3
	MILLI_TS_TO_NANO_TS = TS_TO_NANO_TS / TS_TO_MILLI_TS //10^6
)

func GetNowNanoTS() int64 {
	return time.Now().UnixNano()
}

func TSToNanoTS(ts int64) int64 {
	return ts * TS_TO_NANO_TS
}
