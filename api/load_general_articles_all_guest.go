package api

import (
	"github.com/Ptt-official-app/go-pttbbs/bbs"
	"github.com/gin-gonic/gin"
)

func LoadGeneralArticlesAllGuestWrapper(c *gin.Context) {
	params := NewLoadGeneralArticlesParams()
	path := &LoadGeneralArticlesPath{}
	PathQuery(LoadGeneralArticlesAllGuest, params, path, c)
}

func LoadGeneralArticlesAllGuest(remoteAddr string, params interface{}, path interface{}, c *gin.Context) (result interface{}, statusCode int, err error) {
	userID := bbs.UUserID("guest")

	return LoadGeneralArticles(remoteAddr, userID, params, path, c)
}
