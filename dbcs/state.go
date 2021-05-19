package dbcs

type DBCSState int

const (
	DBCS_STATE_NONE  DBCSState = 0
	DBCS_STATE_LEAD  DBCSState = 1
	DBCS_STATE_TAIL  DBCSState = 2
	DBCS_STATE_COLOR DBCSState = 3
)

func (d DBCSState) String() string {
	switch d {
	case DBCS_STATE_NONE:
		return "none"
	case DBCS_STATE_LEAD:
		return "lead"
	case DBCS_STATE_TAIL:
		return "tail"
	case DBCS_STATE_COLOR:
		return "color"
	default:
		return "unknown"
	}
}
