package api

import (
	"github.com/Ptt-official-app/go-pttbbs/bbs"
	"github.com/gin-gonic/gin"
)

func LoadArticleCommentsAllGuestWrapper(c *gin.Context) {
	params := NewLoadArticleCommentsParams()
	path := &LoadArticleCommentsPath{}
	PathQuery(LoadArticleCommentsAllGuest, params, path, c)
}

func LoadArticleCommentsAllGuest(remoteAddr string, params interface{}, path interface{}, c *gin.Context) (result interface{}, statusCode int, err error) {
	userID := bbs.UUserID("guest")

	return LoadArticleComments(remoteAddr, userID, params, path, c)
}
