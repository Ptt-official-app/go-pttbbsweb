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
	filename := filepath.Join(types.STATIC_DIR, "home.html")

	maxAge := htmlMaxAge()
	processCSRFContent(filename, maxAge, c)
}

func UserHtmlWrapper(c *gin.Context) {
	filename := filepath.Join(types.STATIC_DIR, "user-info.html")

	maxAge := htmlMaxAge()
	processCSRFContent(filename, maxAge, c)
}

func UserChangePasswdHtmlWrapper(c *gin.Context) {
	filename := filepath.Join(types.STATIC_DIR, "change-passwd.html")

	maxAge := htmlMaxAge()
	processCSRFContent(filename, maxAge, c)
}

func htmlMaxAge() (maxAge int) {
	maxAge = HTML_CACHE_CONTROL_TS
	if maxAge < types.CSRF_TOKEN_TS/4 {
		maxAge = types.CSRF_TOKEN_TS / 4
	}

	return maxAge
}
