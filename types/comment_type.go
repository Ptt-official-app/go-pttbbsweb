package types

type CommentType uint8

const (
	COMMENT_TYPE_UNKNOWN   CommentType = 0
	COMMENT_TYPE_RECOMMEND CommentType = 1
	COMMENT_TYPE_BOO       CommentType = 2
	COMMENT_TYPE_COMMENT   CommentType = 3
	COMMENT_TYPE_FORWARD   CommentType = 4

	COMMENT_TYPE_REPLY   CommentType = 5
	COMMENT_TYPE_EDIT    CommentType = 6
	COMMENT_TYPE_DELETED CommentType = 7

	COMMENT_TYPE_BASIC CommentType = 4
)

func (c CommentType) String() string {
	switch c {
	case COMMENT_TYPE_UNKNOWN:
		return "unknown"
	case COMMENT_TYPE_RECOMMEND:
		return "recommend"
	case COMMENT_TYPE_BOO:
		return "boo"
	case COMMENT_TYPE_COMMENT:
		return "comment"
	case COMMENT_TYPE_REPLY:
		return "reply"
	case COMMENT_TYPE_EDIT:
		return "edit"
	case COMMENT_TYPE_FORWARD:
		return "forward"
	case COMMENT_TYPE_DELETED:
		return "deleted"
	default:
		return "unknown"
	}
}
