package api

import (
	"github.com/gin-gonic/gin"
)

func GetUserIDAllGuestWrapper(c *gin.Context) {
	Query(GetUserIDAllGuest, nil, c)
}

func GetUserIDAllGuest(remoteAddr string, user *UserInfo, params interface{}, c *gin.Context) (result interface{}, statusCode int, err error) {
	return GetUserID(remoteAddr, user, params, c)
}
