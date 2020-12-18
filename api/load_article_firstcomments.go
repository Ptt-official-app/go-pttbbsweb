package api

import (
	"github.com/Ptt-official-app/go-openbbsmiddleware/apitypes"
	"github.com/Ptt-official-app/go-openbbsmiddleware/mock"
	"github.com/Ptt-official-app/go-pttbbs/bbs"
	"github.com/gin-gonic/gin"
)

const LOAD_ARTICLE_FIRSTCOMMENTS_R = "/board/:bid/article/:aid/comments/first"

type LoadArticleFirstCommentsParams struct {
}

type LoadArticleFirstCommentsPath struct {
	BBoardID  bbs.BBoardID  `json:"bid"`
	ArticleID bbs.ArticleID `json:"aid"`
}

type LoadArticleFirstCommentsResult struct {
	List []*apitypes.Comment `json:"list"`
}

func NewLoadArticleFirstCommentsParams() *LoadArticleFirstCommentsParams {
	return &LoadArticleFirstCommentsParams{}
}

func LoadArticleFirstCommentsWrapper(c *gin.Context) {
	params := NewLoadArticleFirstCommentsParams()
	path := &LoadArticleFirstCommentsPath{}
	LoginRequiredPathQuery(LoadArticleFirstComments, params, path, c)
}

func LoadArticleFirstComments(remoteAddr string, userID bbs.UUserID, params interface{}, path interface{}, c *gin.Context) (result interface{}, statusCode int, err error) {

	result = mock.CommentListResult
	return result, 200, nil
}
