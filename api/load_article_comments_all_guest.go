package api

import (
	pttbbsapi "github.com/Ptt-official-app/go-pttbbs/api"
	"github.com/Ptt-official-app/go-pttbbs/bbs"
	"github.com/gin-gonic/gin"
)

func LoadArticleCommentsAllGuestWrapper(c *gin.Context) {
	params := NewLoadArticleCommentsParams()
	path := &LoadArticleCommentsPath{}
	PathQuery(LoadArticleCommentsAllGuest, params, path, c)
}

func LoadArticleCommentsAllGuest(remoteAddr string, user *UserInfo, params interface{}, path interface{}, c *gin.Context) (result interface{}, statusCode int, err error) {
	user.UserID = bbs.UUserID(pttbbsapi.GUEST)

	return LoadArticleComments(remoteAddr, user, params, path, c)
}
