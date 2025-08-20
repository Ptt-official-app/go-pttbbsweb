package api

import (
	pttbbsapi "github.com/Ptt-official-app/go-pttbbs/api"
	"github.com/Ptt-official-app/go-pttbbs/bbs"
	"github.com/gin-gonic/gin"
)

func GetArticleDetailAllGuestWrapper(c *gin.Context) {
	params := &GetArticleDetailParams{}
	path := &GetArticleDetailPath{}
	PathQuery(GetArticleDetailAllGuest, params, path, c)
}

func GetArticleDetailAllGuest(remoteAddr string, user *UserInfo, params interface{}, path interface{}, c *gin.Context) (result interface{}, statusCode int, err error) {
	user.UserID = bbs.UUserID(pttbbsapi.GUEST)

	return GetArticleDetail(remoteAddr, user, params, path, c)
}
