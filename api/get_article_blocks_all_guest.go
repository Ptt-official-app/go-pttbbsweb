package api

import (
	"github.com/Ptt-official-app/go-pttbbs/bbs"
	"github.com/gin-gonic/gin"
)

func GetArticleBlocksAllGuestWrapper(c *gin.Context) {
	params := NewGetArticleBlocksParams()
	path := &GetArticleBlocksPath{}
	PathQuery(GetArticleBlocksAllGuest, params, path, c)
}

func GetArticleBlocksAllGuest(remoteAddr string, params interface{}, path interface{}, c *gin.Context) (result interface{}, statusCode int, err error) {
	userID := bbs.UUserID("guest")

	return GetArticleBlocks(remoteAddr, userID, params, path, c)
}
