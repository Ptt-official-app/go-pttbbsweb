package dbcs

type DBCSState int

const (
	DBCS_STATE_NONE  DBCSState = 0
	DBCS_STATE_LEAD  DBCSState = 1
	DBCS_STATE_TAIL  DBCSState = 2
	DBCS_STATE_COLOR DBCSState = 3
)
