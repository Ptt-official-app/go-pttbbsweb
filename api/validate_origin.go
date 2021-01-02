package api

import (
	"net/url"

	"github.com/Ptt-official-app/go-openbbsmiddleware/types"
	"github.com/gin-gonic/gin"
)

//isValidOriginReferer
//
//origin: have origin when CORS or POST.
//referer: <a href> or ajax.
//
//need to check origin and referer when they are not empty strings.
//origin as white-list, referer as black-list for now.
func isValidOriginReferer(c *gin.Context) bool {
	origin := c.GetHeader("Origin")
	if origin == "" {
		origin = c.GetHeader("origin")
	}

	referer := c.GetHeader("Referer")
	if referer == "" {
		referer = c.GetHeader("referer")
	}

	if origin != "" && !isValidOrigin(origin) {
		return false
	}

	if referer != "" && !isValidReferer(referer) {
		return false
	}

	return true
}

func isValidOrigin(origin string) bool {
	if len(types.ALLOW_ORIGINS) == 0 {
		return true
	}

	hostname := getHostnameFromURL(origin)
	isValid, ok := types.ALLOW_ORIGINS_MAP[hostname]
	if !ok || !isValid {
		return false
	}

	return true
}

func isValidReferer(referer string) bool {
	if len(types.BLOCKED_REFERERS) == 0 {
		return true
	}

	hostname := getHostnameFromURL(referer)
	isBlocked, ok := types.BLOCKED_REFERERS_MAP[hostname]
	if !ok || isBlocked {
		return false
	}

	return true
}

func getHostnameFromURL(referer string) (hostname string) {
	u, err := url.Parse(referer)
	if err != nil {
		return referer
	}
	return u.Host
}
