package api

import (
	"strings"

	"github.com/gin-gonic/gin"
)

func JSON(theFunc ApiFunc, params interface{}, c *gin.Context) {
	err := c.ShouldBindJSON(params)
	if err != nil {
		processResult(c, nil, 400, err)
		return
	}

	process(theFunc, params, c)
}

func Query(theFunc ApiFunc, params interface{}, c *gin.Context) {
	err := c.ShouldBindQuery(params)
	if err != nil {
		processResult(c, nil, 400, err)
		return
	}

	process(theFunc, params, c)
}

func process(theFunc ApiFunc, params interface{}, c *gin.Context) {

	remoteAddr := strings.TrimSpace(c.ClientIP())
	if !isValidRemoteAddr(remoteAddr) {
		processResult(c, nil, 400, ErrInvalidRemoteAddr)
		return
	}

	result, statusCode, err := theFunc(remoteAddr, params, c)
	processResult(c, result, statusCode, err)
}
