package api

import (
	"github.com/Ptt-official-app/go-openbbsmiddleware/apitypes"
	"github.com/Ptt-official-app/go-openbbsmiddleware/mock"
	"github.com/Ptt-official-app/go-pttbbs/bbs"
	"github.com/gin-gonic/gin"
)

const LOAD_USER_COMMENTS_R = "/user/:user_id/comments"

type LoadUserCommentsParams struct {
	StartIdx   string `json:"start_idx,omitempty" form:"start_idx,omitempty" url:"start_idx,omitempty"`
	Descending bool   `json:"desc,omitempty"  form:"desc,omitempty" url:"desc,omitempty"`
	Max        int    `json:"max,omitempty" form:"max,omitempty" url:"max,omitempty"`
}

type LoadUserCommentsPath struct {
	UserID bbs.UUserID `json:"user_id"`
}

type LoadUserCommentsResult struct {
	List    []*apitypes.Comment `json:"list"`
	NextIdx string              `json:"next_idx"`
}

func NewLoadUserCommentsParams() *LoadUserCommentsParams {
	return &LoadUserCommentsParams{
		Descending: DEFAULT_DESCENDING,
		Max:        DEFAULT_MAX_LIST,
	}
}

func LoadUserCommentsWrapper(c *gin.Context) {
	params := NewLoadUserCommentsParams()
	path := &LoadUserCommentsPath{}
	LoginRequiredPathQuery(LoadUserComments, params, path, c)
}

func LoadUserComments(remoteAddr string, userID bbs.UUserID, params interface{}, path interface{}, c *gin.Context) (result interface{}, statusCode int, err error) {

	result = mock.CommentListResult
	return result, 200, nil
}
