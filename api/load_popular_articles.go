package api

import (
	"github.com/Ptt-official-app/go-openbbsmiddleware/mock"
	"github.com/Ptt-official-app/go-openbbsmiddleware/types"
	"github.com/gin-gonic/gin"
)

const LOAD_POPULAR_ARTICLES_R = "/articles/popular"

type LoadPopularArticlesParams struct{}

type LoadPopularArticlesResult struct {
	List    []*types.ArticleSummary `json:"list"`
	NextIdx string                  `json:"next_idx"`
}

func LoadPopularArticles(remoteAddr string, userID string, params interface{}, c *gin.Context) (result interface{}, statusCode int, err error) {

	result = mock.ArticleListResult
	return result, 200, nil
}
