package api

import (
	"time"

	"github.com/Ptt-official-app/go-openbbsmiddleware/types"
	"github.com/gin-gonic/gin"
)

func setTokenToCookie(c *gin.Context, accessToken string) {
	if c == nil {
		return
	}

	expires := time.Now().Add(3 * 86400 * time.Second)

	expiresStr := expires.Format("")
	c.Header("Set-Cookie", "token="+accessToken+";Expires="+expiresStr+";SameSite=strict;"+types.ACCESS_TOKEN_COOKIE_SUFFIX)
}
