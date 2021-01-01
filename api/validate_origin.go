package api

import (
	"github.com/Ptt-official-app/go-openbbsmiddleware/types"
	"github.com/gin-gonic/gin"
)

func isValidOrigin(c *gin.Context) bool {
	origin := c.GetHeader("Origin")

	if len(types.ALLOW_ORIGINS_MAP) == 0 {
		return true
	}

	isValid, ok := types.ALLOW_ORIGINS_MAP[origin]
	if ok && isValid {
		return true
	}

	return false
}
