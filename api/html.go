package api

import (
	"path/filepath"

	"github.com/Ptt-official-app/go-openbbsmiddleware/types"
	"github.com/gin-gonic/gin"
)

type HTMLPath struct {
	Path string `uri:"path"`
}

const INDEX_HTML_R = "/index.html"

func IndexHtmlWrapper(c *gin.Context) {
	filename := filepath.Join(types.STATIC_DIR, "home.html")

	maxAge := htmlMaxAge()
	processCSRFContent(filename, maxAge, c)
}

const USER_HTML_R = "/user/:user_id"

func UserHtmlWrapper(c *gin.Context) {
	filename := filepath.Join(types.STATIC_DIR, "user-info.html")

	maxAge := htmlMaxAge()
	processCSRFContent(filename, maxAge, c)
}

const USER_CHANGE_PASSWD_HTML_R = "/user/:user_id/resetpassword"

func UserChangePasswdHtmlWrapper(c *gin.Context) {
	filename := filepath.Join(types.STATIC_DIR, "change-passwd.html")

	maxAge := htmlMaxAge()
	processCSRFContent(filename, maxAge, c)
}

const USER_CHANGE_EMAIL_HTML_R = "/user/:user_id/changeemail"

func UserChangeEmailHtmlWrapper(c *gin.Context) {
	filename := filepath.Join(types.STATIC_DIR, "change-email.html")

	maxAge := htmlMaxAge()
	processCSRFContent(filename, maxAge, c)
}

const USER_SET_ID_EMAIL_HTML_R = "/user/:user_id/setidemail"

func UserSetIDEmailHtmlWrapper(c *gin.Context) {
	filename := filepath.Join(types.STATIC_DIR, "set-id-email.html")

	maxAge := htmlMaxAge()
	processCSRFContent(filename, maxAge, c)
}

func htmlMaxAge() (maxAge int) {
	maxAge = HTML_CACHE_CONTROL_TS

	return maxAge
}
