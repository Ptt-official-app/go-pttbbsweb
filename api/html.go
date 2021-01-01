package api

import (
	"path/filepath"

	"github.com/Ptt-official-app/go-openbbsmiddleware/types"
	"github.com/gin-gonic/gin"
)

type HTMLPath struct {
	Path string `uri:"path"`
}

func IndexHtmlWrapper(c *gin.Context) {
	filename := filepath.Join(types.STATIC_DIR, "index.html")
	processCSRFContent(filename, c)
}
