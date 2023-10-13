package api

import (
	"strings"

	"github.com/gin-gonic/gin"
)

func RedirectPathQuery(theFunc RedirectPathAPIFunc, params interface{}, path interface{}, c *gin.Context) {
	err := c.ShouldBindQuery(params)
	if err != nil {
		processResult(c, nil, 400, err, "")
		return
	}

	redirectPathProcess(theFunc, params, path, c)
}

func redirectPathProcess(theFunc RedirectPathAPIFunc, params interface{}, path interface{}, c *gin.Context) {
	err := c.ShouldBindUri(path)
	if err != nil {
		processResult(c, nil, 400, err, "")
		return
	}

	remoteAddr := strings.TrimSpace(c.ClientIP())
	if !isValidRemoteAddr(remoteAddr) {
		processResult(c, nil, 403, ErrInvalidRemoteAddr, "")
		return
	}

	if !isValidOriginReferer(c) {
		processResult(c, nil, 403, ErrInvalidOrigin, "")
		return
	}

	redirectPath, statusCode := theFunc(remoteAddr, params, path, c)
	processRedirectResult(c, redirectPath, statusCode)
}
