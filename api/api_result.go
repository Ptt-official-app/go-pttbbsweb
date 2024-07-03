package api

import (
	"github.com/Ptt-official-app/go-pttbbs/bbs"
	"github.com/Ptt-official-app/go-pttbbsweb/types"
	"github.com/gin-gonic/gin"
)

func processResult(c *gin.Context, result interface{}, statusCode int, err error, userID bbs.UUserID) {
	setHeader(c)

	if statusCode == 200 {
		c.JSON(200, result)
	} else {
		c.JSON(statusCode, &errResult{Msg: err.Error(), TokenUser: userID})
	}
}

func processStringResult(c *gin.Context, content string, contentType string) {
	setHeader(c)

	c.Header("Content-Type", contentType)
	c.String(200, "%v", content)
}

func processRedirectResult(c *gin.Context, redirectPath string, statusCode int) {
	setHeader(c)

	c.Redirect(statusCode, redirectPath)
}

func setHeader(c *gin.Context) {
	if !types.IS_ALLOW_CROSSDOMAIN {
		return
	}

	origin := c.GetHeader("Origin")

	if origin == "" {
		return
	}

	requestHeaders := c.GetHeader("Access-Control-Request-Headers")

	c.Header("X-Frame-Options", "SAMEORIGIN")
	c.Header("Access-Control-Allow-Methods", "GET,POST,OPTIONS")
	c.Header("Access-Control-Allow-Credentials", "true")
	c.Header("Access-Control-Allow-Origin", origin)
	if requestHeaders != "" {
		c.Header("Access-Control-Allow-Headers", requestHeaders)
	}
}
