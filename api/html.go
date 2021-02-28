package api

import (
	"path/filepath"

	"github.com/Ptt-official-app/go-openbbsmiddleware/types"
	"github.com/gin-gonic/gin"
)

type HTMLPath struct {
	Path string `uri:"path"`
}

const ROOT_HTML_R = "/"
const INDEX_HTML_R = "/index.html"

func IndexHtmlWrapper(c *gin.Context) {
	filename := filepath.Join(types.STATIC_DIR, "home.html")

	maxAge := htmlMaxAge()
	processCSRFContent(filename, maxAge, c)
}

const BOARDS_HTML_R = "/boards"

const BOARDS_POPULAR_HTML_R = "/boards/popular"

const BOARDS_FAVORITES_HTML_R = "/user/:user_id/favorites"

const ARTICLES_HTML_R = "/board/:bid/articles"

const ARTICLE_HTML_R = "/board/:bid/article/:aid"

func AllHtmlWrapper(c *gin.Context) {
	filename := filepath.Join(types.STATIC_DIR, "all.html")

	maxAge := htmlMaxAge()
	processCSRFContent(filename, maxAge, c)
}

const LOGIN_HTML_R = "/login"

func LoginHtmlWrapper(c *gin.Context) {
	filename := filepath.Join(types.STATIC_DIR, "login.html")

	maxAge := htmlMaxAge()
	processCSRFContent(filename, maxAge, c)
}

const REGISTER_HTML_R = "/register"

func RegisterHtmlWrapper(c *gin.Context) {
	filename := filepath.Join(types.STATIC_DIR, "register.html")

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

const USER_ATTEMPT_CHANGE_EMAIL_HTML_R = "/user/:user_id/attemptchangeemail"

func UserAttemptChangeEmailHtmlWrapper(c *gin.Context) {
	filename := filepath.Join(types.STATIC_DIR, "attempt-change-email.html")

	maxAge := htmlMaxAge()
	processCSRFContent(filename, maxAge, c)
}

const USER_CHANGE_EMAIL_HTML_R = "/user/:user_id/changeemail"

func UserChangeEmailHtmlWrapper(c *gin.Context) {
	filename := filepath.Join(types.STATIC_DIR, "change-email.html")

	maxAge := htmlMaxAge()
	processCSRFContent(filename, maxAge, c)
}

const USER_ATTEMPT_SET_ID_EMAIL_HTML_R = "/user/:user_id/attemptsetidemail"

func UserAttemptSetIDEmailHtmlWrapper(c *gin.Context) {
	filename := filepath.Join(types.STATIC_DIR, "attempt-set-id-email.html")

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
