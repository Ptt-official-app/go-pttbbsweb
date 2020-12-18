package api

import "github.com/gin-gonic/gin"

const INDEX_R = "/"

type IndexParams struct {
	In int `form:"in,omitempty"`
}

func NewIndexParams() *IndexParams {
	return &IndexParams{}
}

type IndexResult struct {
	Data interface{}
}

func IndexWrapper(c *gin.Context) {
	params := NewIndexParams()
	Query(Index, params, c)
}

func Index(remoteAddr string, params interface{}, c *gin.Context) (result interface{}, statusCode int, err error) {
	result = &IndexResult{Data: params}
	return result, 200, nil
}
