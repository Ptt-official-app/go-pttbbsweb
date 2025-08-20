package api

import (
	pttbbsapi "github.com/Ptt-official-app/go-pttbbs/api"
	"github.com/Ptt-official-app/go-pttbbs/bbs"
	"github.com/gin-gonic/gin"
)

func GetArticleBlocksAllGuestWrapper(c *gin.Context) {
	params := NewGetArticleBlocksParams()
	path := &GetArticleBlocksPath{}
	PathQuery(GetArticleBlocksAllGuest, params, path, c)
}

func GetArticleBlocksAllGuest(remoteAddr string, user *UserInfo, params interface{}, path interface{}, c *gin.Context) (result interface{}, statusCode int, err error) {
	user.UserID = bbs.UUserID(pttbbsapi.GUEST)
	return GetArticleBlocks(remoteAddr, user, params, path, c)
}
