package api

import (
	"github.com/gin-gonic/gin"
)

func processResult(c *gin.Context, result interface{}, statusCode int, err error) {
	if statusCode == 200 {
		c.JSON(200, result)
	} else {
		c.JSON(statusCode, &errResult{err.Error()})
	}
}
