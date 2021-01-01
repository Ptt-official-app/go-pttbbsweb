package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func OptionsWrapper(c *gin.Context) {
	processResult(c, struct{}{}, http.StatusOK, nil)
}
