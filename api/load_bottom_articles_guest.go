package api

import (
	"github.com/Ptt-official-app/go-pttbbs/bbs"
	"github.com/gin-gonic/gin"
)

func LoadBottomArticlesAllGuestWrapper(c *gin.Context) {
	path := &LoadBottomArticlesPath{}
	PathQuery(LoadBottomArticlesAllGuest, nil, path, c)
}

func LoadBottomArticlesAllGuest(remoteAddr string, params interface{}, path interface{}, c *gin.Context) (result interface{}, statusCode int, err error) {
	userID := bbs.UUserID("guest")

	return LoadBottomArticles(remoteAddr, userID, params, path, c)
}
