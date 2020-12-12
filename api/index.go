package api

import "github.com/gin-gonic/gin"

type IndexParams struct {
	In int `form:"in,omitempty"`
}

type IndexResult struct {
	Data interface{}
}

func Index(remoteAddr string, params interface{}, c *gin.Context) (result interface{}, statusCode int, err error) {
	result = &IndexResult{Data: params}
	return result, 200, nil
}
