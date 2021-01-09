package api

import (
	"github.com/Ptt-official-app/go-openbbsmiddleware/types"
	"github.com/Ptt-official-app/go-pttbbs/bbs"
	"github.com/gin-gonic/gin"
)

func setTokenToCookie(c *gin.Context, accessToken string) {
	setCookie(c, types.ACCESS_TOKEN_NAME, accessToken, types.ACCESS_TOKEN_EXPIRE_TS_DURATION, true)
}

func loginRegisterRedirectUrl(userID bbs.UUserID, c *gin.Context) string {
	return "/user/" + string(userID)
}
