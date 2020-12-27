package dbcs

type DBCSState int

const (
	DBCS_STATE_NONE DBCSState = iota
	DBCS_STATE_LEAD
	DBCS_STATE_TAIL
	DBCS_STATE_COLOR
)
