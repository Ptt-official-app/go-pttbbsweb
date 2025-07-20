package api

import (
	pttbbsapi "github.com/Ptt-official-app/go-pttbbs/api"
	"github.com/Ptt-official-app/go-pttbbs/bbs"
	"github.com/gin-gonic/gin"
)

func GetUserIDAllGuestWrapper(c *gin.Context) {
	Query(GetUserIDAllGuest, nil, c)
}

func GetUserIDAllGuest(remoteAddr string, params interface{}, c *gin.Context) (result interface{}, statusCode int, err error) {
	return &GetUserIDResult{
		UserID: bbs.UUserID(pttbbsapi.GUEST),

		TokenUser: bbs.UUserID(pttbbsapi.GUEST),
	}, 200, nil
}
