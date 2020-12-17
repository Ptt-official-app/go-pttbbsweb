package api

import (
	"github.com/Ptt-official-app/go-openbbsmiddleware/mock"
	"github.com/Ptt-official-app/go-openbbsmiddleware/types"
	"github.com/Ptt-official-app/go-pttbbs/bbs"
	"github.com/gin-gonic/gin"
)

const LOAD_ARTICLE_COMMENTS_R = "/board/:bid/article/:aid/comments"

type LoadArticleCommentsParams struct {
	StartIdx string `json:"start_idx,omitempty" form:"start_idx,omitempty" url:"start_idx,omitempty"`
	Max      int    `json:"max,omitempty" form:"max,omitempty" url:"max,omitempty"`
}

type LoadArticleCommentsPath struct {
	BBoardID  bbs.BBoardID  `json:"bid"`
	ArticleID bbs.ArticleID `json:"aid"`
}

type LoadArticleCommentsResult struct {
	List    []*types.Comment `json:"list"`
	NextIdx string           `json:"next_idx"`
}

func LoadArticleComments(remoteAddr string, userID bbs.UUserID, params interface{}, path interface{}, c *gin.Context) (result interface{}, statusCode int, err error) {

	result = mock.CommentListResult
	return result, 200, nil
}
