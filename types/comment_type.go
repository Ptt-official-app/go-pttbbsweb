package types

type CommentType uint8

const (
	_ CommentType = iota
	COMMENT_TYPE_RECOMMEND
	COMMENT_TYPE_BOO
	COMMENT_TYPE_COMMENT
	COMMENT_TYPE_REPLY
)
