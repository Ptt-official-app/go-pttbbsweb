package dbcs

import "time"

const (
	N_FIRST_COMMENTS      = 10
	COMMENT_STEP_DURATION = 1 * time.Millisecond
	REPLY_STEP_NANO_TS    = 100000 //0.1 millisecond

	DEFAULT_LINE_BYTES = 200

	LEN_OLD_RECOMMEND_DATE = 5
	LEN_RECOMMEND_DATE     = 11
)
