package api

import (
	"github.com/Ptt-official-app/go-openbbsmiddleware/mock"
	"github.com/Ptt-official-app/go-openbbsmiddleware/types"
	"github.com/gin-gonic/gin"
)

const LOAD_USER_COMMENTS_R = "/user/:user_id/comments"

type LoadUserCommentsParams struct {
	StartIdx string `json:"start_idx,omitempty" form:"start_idx,omitempty" url:"start_idx,omitempty"`
	Max      int    `json:"max,omitempty" form:"max,omitempty" url:"max,omitempty"`
}

type LoadUserCommentsPath struct {
	UserID string `json:"user_id"`
}

type LoadUserCommentsResult struct {
	List    []*types.Comment `json:"list"`
	NextIdx string           `json:"next_idx"`
}

func LoadUserComments(remoteAddr string, userID string, params interface{}, path interface{}, c *gin.Context) (result interface{}, statusCode int, err error) {

	result = mock.CommentListResult
	return result, 200, nil
}
