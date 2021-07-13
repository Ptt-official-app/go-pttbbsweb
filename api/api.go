package api

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

func JSON(theFunc ApiFunc, params interface{}, c *gin.Context) {
	err := c.ShouldBindJSON(params)
	if err != nil {
		processResult(c, nil, 400, err)
		return
	}

	process(theFunc, params, c)
}

func Form(theFunc ApiFunc, params interface{}, c *gin.Context) {
	err := c.ShouldBindWith(params, binding.Form)
	if err != nil {
		processResult(c, nil, 400, err)
		return
	}

	process(theFunc, params, c)
}

func FormJSON(theFunc ApiFunc, params interface{}, c *gin.Context) {
	err := c.ShouldBindJSON(params)
	if err != nil {
		err = c.ShouldBindWith(params, binding.Form)
		if err != nil {
			processResult(c, nil, 400, err)
			return
		}
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

	if !isValidOriginReferer(c) {
		processResult(c, nil, 403, ErrInvalidOrigin)
		return
	}

	result, statusCode, err := theFunc(remoteAddr, params, c)
	processResult(c, result, statusCode, err)
}
