package api

import (
	"github.com/Ptt-official-app/go-openbbsmiddleware/apitypes"
	"github.com/Ptt-official-app/go-openbbsmiddleware/mock"
	"github.com/Ptt-official-app/go-pttbbs/bbs"
	"github.com/gin-gonic/gin"
)

const LOAD_POPULAR_ARTICLES_R = "/articles/popular"

type LoadPopularArticlesParams struct {
	StartIdx   string `json:"start_idx,omitempty" form:"start_idx,omitempty" url:"start_idx,omitempty"`
	Max        int    `json:"limit,omitempty" form:"limit,omitempty" url:"limit,omitempty"`
	Descending bool   `json:"desc,omitempty"  form:"desc,omitempty" url:"desc,omitempty"`
}

type LoadPopularArticlesResult struct {
	List    []*apitypes.ArticleSummary `json:"list"`
	NextIdx string                     `json:"next_idx"`

	TokenUser bbs.UUserID `json:"tokenuser,omitempty"`
}

func NewLoadPopularArticlesParams() *LoadPopularArticlesParams {
	return &LoadPopularArticlesParams{
		Max:        DEFAULT_MAX_LIST,
		Descending: DEFAULT_DESCENDING,
	}
}

func LoadPopularArticlesWrapper(c *gin.Context) {
	params := NewLoadPopularArticlesParams()
	LoginRequiredQuery(LoadPopularArticles, params, c)
}

func LoadPopularArticles(remoteAddr string, userID bbs.UUserID, params interface{}, c *gin.Context) (result interface{}, statusCode int, err error) {
	result = mock.ArticleListResult
	return result, 200, nil
}
