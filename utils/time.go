package utils

import "time"

const (
	TS_TO_NANO_TS = int64(1000000000) //10^9
)

func GetNowNanoTS() int64 {
	return time.Now().UnixNano()
}

func TSToNanoTS(ts int64) int64 {
	return ts * TS_TO_NANO_TS
}
