package api

import (
	pttbbsapi "github.com/Ptt-official-app/go-pttbbs/api"
	"github.com/Ptt-official-app/go-pttbbs/bbs"
	"github.com/gin-gonic/gin"
)

func LoadBottomArticlesAllGuestWrapper(c *gin.Context) {
	path := &LoadBottomArticlesPath{}
	PathQuery(LoadBottomArticlesAllGuest, nil, path, c)
}

func LoadBottomArticlesAllGuest(remoteAddr string, user *UserInfo, params interface{}, path interface{}, c *gin.Context) (result interface{}, statusCode int, err error) {
	user.UserID = bbs.UUserID(pttbbsapi.GUEST)

	return LoadBottomArticles(remoteAddr, user, params, path, c)
}
