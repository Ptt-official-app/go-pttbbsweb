package api

import (
	"github.com/Ptt-official-app/go-openbbsmiddleware/mock"
	"github.com/Ptt-official-app/go-openbbsmiddleware/types"
	"github.com/Ptt-official-app/go-pttbbs/bbs"
	"github.com/gin-gonic/gin"
)

const LOAD_BOTTOM_ARTICLES_R = "/board/:bid/articles/bottom"

type LoadBottomArticlesParams struct{}

type LoadBottomArticlesPath struct {
	BBoardID bbs.BBoardID `json:"bid"`
}

type LoadBottomArticlesResult struct {
	List    []*types.ArticleSummary `json:"list"`
	NextIdx string                  `json:"next_idx"`
}

func LoadBottomArticles(remoteAddr string, userID bbs.UUserID, params interface{}, path interface{}, c *gin.Context) (result interface{}, statusCode int, err error) {

	result = mock.ArticleListResult
	return result, 200, nil
}
