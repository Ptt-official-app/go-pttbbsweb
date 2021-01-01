package api

import (
	"path/filepath"

	"github.com/Ptt-official-app/go-openbbsmiddleware/types"
	"github.com/gin-gonic/gin"
)

type JSPath struct {
	Path string `uri:"path"`
}

func JSWrapper(c *gin.Context) {

	path := &JSPath{}
	err := c.ShouldBindUri(path)
	if err != nil {
		processResult(c, nil, 400, err)
		return
	}

	filename := filepath.Join(types.STATIC_DIR, "static", "js", path.Path)
	processCSRFContent(filename, c)
}
