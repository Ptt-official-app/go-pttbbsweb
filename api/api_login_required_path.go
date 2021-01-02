package api

import (
	"strings"

	"github.com/gin-gonic/gin"
)

func LoginRequiredPathJSON(theFunc LoginRequiredPathApiFunc, params interface{}, path interface{}, c *gin.Context) {
	err := c.ShouldBindJSON(params)
	if err != nil {
		processResult(c, nil, 400, err)
		return
	}

	loginRequiredPathProcess(theFunc, params, path, c)
}

func LoginRequiredPathQuery(theFunc LoginRequiredPathApiFunc, params interface{}, path interface{}, c *gin.Context) {
	err := c.ShouldBindQuery(params)
	if err != nil {
		processResult(c, nil, 400, err)
		return
	}

	loginRequiredPathProcess(theFunc, params, path, c)
}

func loginRequiredPathProcess(theFunc LoginRequiredPathApiFunc, params interface{}, path interface{}, c *gin.Context) {

	err := c.ShouldBindUri(path)
	if err != nil {
		processResult(c, nil, 400, err)
		return
	}

	remoteAddr := strings.TrimSpace(c.ClientIP())
	if !isValidRemoteAddr(remoteAddr) {
		processResult(c, nil, 403, ErrInvalidRemoteAddr)
		return
	}

	if !isValidOriginReferer(c) {
		processResult(c, nil, 403, ErrInvalidOrigin)
		return
	}
	userID, err := verifyJwt(c)
	if err != nil {
		processResult(c, nil, 401, err)
		return
	}

	result, statusCode, err := theFunc(remoteAddr, userID, params, path, c)
	processResult(c, result, statusCode, err)

}
