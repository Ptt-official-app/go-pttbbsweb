package api

import (
	"github.com/gin-gonic/gin"
)

type ApiFunc func(remoteAddr string, params interface{}, c *gin.Context) (result interface{}, statusCode int, err error)

type LoginRequiredApiFunc func(remoteAddr string, userID string, params interface{}, c *gin.Context) (result interface{}, statusCode int, err error)
