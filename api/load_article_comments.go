package api

import (
	"github.com/Ptt-official-app/go-openbbsmiddleware/apitypes"
	"github.com/Ptt-official-app/go-openbbsmiddleware/mock"
	"github.com/Ptt-official-app/go-pttbbs/bbs"
	"github.com/gin-gonic/gin"
)

const LOAD_ARTICLE_COMMENTS_R = "/board/:bid/article/:aid/comments"

type LoadArticleCommentsParams struct {
	StartIdx   string `json:"start_idx,omitempty" form:"start_idx,omitempty" url:"start_idx,omitempty"`
	Descending bool   `json:"desc,omitempty"  form:"desc,omitempty" url:"desc,omitempty"`
	Max        int    `json:"limit,omitempty" form:"limit,omitempty" url:"limit,omitempty"`
}

type LoadArticleCommentsPath struct {
	BBoardID  bbs.BBoardID  `json:"bid"`
	ArticleID bbs.ArticleID `json:"aid"`
}

type LoadArticleCommentsResult struct {
	List    []*apitypes.Comment `json:"list"`
	NextIdx string              `json:"next_idx"`
}

func NewLoadArticleCommentsParams() *LoadArticleCommentsParams {
	return &LoadArticleCommentsParams{
		Descending: DEFAULT_DESCENDING,
		Max:        DEFAULT_MAX_LIST,
	}
}

func LoadArticleCommentsWrapper(c *gin.Context) {
	params := NewLoadArticleCommentsParams()
	path := &LoadArticleCommentsPath{}
	LoginRequiredPathQuery(LoadArticleComments, params, path, c)
}

func LoadArticleComments(remoteAddr string, userID bbs.UUserID, params interface{}, path interface{}, c *gin.Context) (result interface{}, statusCode int, err error) {

	result = mock.CommentListResult
	return result, 200, nil
}
