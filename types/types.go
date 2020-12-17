package types

type CommentID string

const (
	TS_TO_NANO_TS = NanoTS(1000000000) //10^9
)

type Time8 int64

type NanoTS int64

func (t Time8) ToNanoTS() NanoTS {
	return NanoTS(t) * TS_TO_NANO_TS
}

func (t NanoTS) ToTS() Time8 {
	return Time8(t / TS_TO_NANO_TS)
}
