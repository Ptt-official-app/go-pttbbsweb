package dbcs

import "github.com/Ptt-official-app/pttbbs-backend/types"

var (
	COMMENT_STEP_DIFF_NANO_TS  types.NanoTS = 2 * 60 * types.TS_TO_NANO_TS    // 2 mins
	COMMENT_STEP_DIFF2_NANO_TS types.NanoTS = 2 * 86400 * types.TS_TO_NANO_TS // 2 days
)
