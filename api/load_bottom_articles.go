package api

import (
	"github.com/Ptt-official-app/go-openbbsmiddleware/apitypes"
	"github.com/Ptt-official-app/go-openbbsmiddleware/mock"
	"github.com/Ptt-official-app/go-pttbbs/bbs"
	"github.com/gin-gonic/gin"
)

const LOAD_BOTTOM_ARTICLES_R = "/board/:bid/articles/bottom"

type LoadBottomArticlesParams struct{}

type LoadBottomArticlesPath struct {
	FBoardID apitypes.FBoardID `json:"bid"`
}

type LoadBottomArticlesResult struct {
	List    []*apitypes.ArticleSummary `json:"list"`
	NextIdx string                     `json:"next_idx"`
}

func LoadBottomArticlesWrapper(c *gin.Context) {
	params := &LoadBottomArticlesParams{}
	path := &LoadBottomArticlesPath{}
	LoginRequiredPathQuery(LoadBottomArticles, params, path, c)
}

func LoadBottomArticles(remoteAddr string, userID bbs.UUserID, params interface{}, path interface{}, c *gin.Context) (result interface{}, statusCode int, err error) {

	result = mock.ArticleListResult
	return result, 200, nil
}
