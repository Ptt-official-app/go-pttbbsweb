package api

import (
	"github.com/Ptt-official-app/go-openbbsmiddleware/apitypes"
	"github.com/Ptt-official-app/go-openbbsmiddleware/mock"
	"github.com/Ptt-official-app/go-pttbbs/bbs"
	"github.com/gin-gonic/gin"
)

const LOAD_USER_ARTICLES_R = "/user/:user_id/articles"

type LoadUserArticlesParams struct {
	StartIdx   string `json:"start_idx,omitempty" form:"start_idx,omitempty" url:"start_idx,omitempty"`
	Descending bool   `json:"desc,omitempty"  form:"desc,omitempty" url:"desc,omitempty"`
	Max        int    `json:"limit,omitempty" form:"limit,omitempty" url:"limit,omitempty"`
}

type LoadUserArticlesPath struct {
	UserID bbs.UUserID `json:"user_id"`
}

type LoadUserArticlesResult struct {
	List    []*apitypes.ArticleSummary `json:"list"`
	NextIdx string                     `json:"next_idx"`
}

func NewUserArticlesParams() *LoadUserArticlesParams {
	return &LoadUserArticlesParams{
		Descending: DEFAULT_DESCENDING,
		Max:        DEFAULT_MAX_LIST,
	}
}

func LoadUserArticlesWrapper(c *gin.Context) {
	params := NewUserArticlesParams()
	path := &LoadUserArticlesPath{}
	LoginRequiredPathQuery(LoadUserArticles, params, path, c)
}

func LoadUserArticles(remoteAddr string, userID bbs.UUserID, params interface{}, path interface{}, c *gin.Context) (result interface{}, statusCode int, err error) {

	result = mock.ArticleListResult
	return result, 200, nil
}
