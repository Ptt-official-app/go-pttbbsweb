package api

import (
	"github.com/Ptt-official-app/go-pttbbs/bbs"
	"github.com/gin-gonic/gin"
)

func GetArticleDetailAllGuestWrapper(c *gin.Context) {
	params := &GetArticleDetailParams{}
	path := &GetArticleDetailPath{}
	PathQuery(GetArticleDetailAllGuest, params, path, c)
}

func GetArticleDetailAllGuest(remoteAddr string, params interface{}, path interface{}, c *gin.Context) (result interface{}, statusCode int, err error) {
	userID := bbs.UUserID("guest")

	return GetArticleDetail(remoteAddr, userID, params, path, c)
}
