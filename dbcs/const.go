package dbcs

import "time"

const (
	N_FIRST_COMMENTS      = 10
	COMMENT_STEP_DURATION = 1 * time.Millisecond
	REPLY_STEP_NANO_TS    = 100000 //0.1 millisecond
)
