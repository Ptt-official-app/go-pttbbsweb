package api

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/Ptt-official-app/go-openbbsmiddleware/types"
	"github.com/Ptt-official-app/go-openbbsmiddleware/utils"
	pttbbsapi "github.com/Ptt-official-app/go-pttbbs/api"
	"github.com/Ptt-official-app/go-pttbbs/bbs"
	"github.com/gin-gonic/gin"
	"gopkg.in/square/go-jose.v2"
	"gopkg.in/square/go-jose.v2/jwt"
)

func verifyJwt(c *gin.Context) (userID bbs.UUserID, err error) {
	jwt := pttbbsapi.GetJwt(c) //get jwt from access-token

	if jwt == "" {
		jwt = utils.GetCookie(c, types.ACCESS_TOKEN_NAME)
	}

	userID, clientInfoStr, err := pttbbsapi.VerifyJwt(jwt)
	if err != nil {
		return "", err
	}

	clientInfo := &ClientInfo{}
	err = json.Unmarshal([]byte(clientInfoStr), clientInfo)

	if clientInfo.ClientType == CLIENT_TYPE_APP {
		return userID, nil
	}

	csrfToken := c.GetHeader("X-CSRFToken")
	if len(csrfToken) == 0 {
		return "", ErrInvalidToken
	}

	cookieCSRFToken := utils.GetCookie(c, types.CSRF_TOKEN)
	if cookieCSRFToken == "" {
		return "", ErrInvalidToken
	}

	if csrfToken != cookieCSRFToken {
		return "", ErrInvalidToken
	}

	if !isValidCSRFToken(csrfToken) {
		return "", ErrInvalidToken
	}

	return userID, nil
}

func createCSRFToken() (string, error) {
	var err error

	sig, err := jose.NewSigner(jose.SigningKey{Algorithm: jose.HS256, Key: types.CSRF_SECRET}, (&jose.SignerOptions{}).WithType("JWT"))
	if err != nil {
		return "", err
	}

	cl := &pttbbsapi.JwtClaim{
		Expire: jwt.NewNumericDate(time.Now().Add(types.CSRF_TOKEN_TS_DURATION)),
	}

	raw, err := jwt.Signed(sig).Claims(cl).CompactSerialize()
	if err != nil {
		return "", err
	}

	return raw, nil
}

func isValidCSRFToken(raw string) bool {

	tok, err := jwt.ParseSigned(raw)
	if err != nil {
		return false
	}

	cl := &pttbbsapi.JwtClaim{}
	if err := tok.Claims(types.CSRF_SECRET, cl); err != nil {
		return false
	}

	currentNanoTS := jwt.NewNumericDate(time.Now())
	if *currentNanoTS > *cl.Expire {
		return false
	}

	return true
}

func processCSRFContent(filename string, cacheControlMaxAge int, c *gin.Context) {
	if !isValidOriginReferer(c) {
		processResult(c, nil, 403, ErrInvalidOrigin)
		return
	}

	file, err := os.Open(filename)
	if err != nil {
		processResult(c, nil, 404, ErrFileNotFound)
		return
	}
	defer file.Close()

	reader := io.Reader(file)
	contentBytes, err := ioutil.ReadAll(reader)
	if err != nil {
		processResult(c, nil, 500, ErrInvalidPath)
	}

	ext := filepath.Ext(filename)
	mimeType, _ := MIME_TYPE_MAP[ext]

	content := string(contentBytes)

	csrfToken := utils.GetCookie(c, types.CSRF_TOKEN)
	if csrfToken == "" {
		csrfToken, _ = createCSRFToken()
		setCookie(c, types.CSRF_TOKEN, csrfToken, types.CSRF_TOKEN_TS_DURATION, true)
	}
	content = strings.Replace(content, "__CSRFTOKEN__", csrfToken, 1)

	c.Header("Cache-Control", "max-age="+strconv.Itoa(cacheControlMaxAge))

	processStringResult(c, content, mimeType)
}

func setCookie(c *gin.Context, name string, value string, expireDuration time.Duration, isHttpOnly bool) {
	if c == nil {
		return
	}
	setCookie := name + "=" + value + ";Domain=" + types.COOKIE_DOMAIN + ";Path=/;"
	if expireDuration != 0 {
		expires := time.Now().Add(expireDuration)
		expiresStr := expires.Format("Mon, Jan 2 2006 15:04:05 MST")
		setCookie += "Expires=" + expiresStr + ";"
	}
	if isHttpOnly {
		setCookie += "HttpOnly;"
	}

	setCookie += "SameSite=Lax;" + types.TOKEN_COOKIE_SUFFIX
	c.Header("Set-Cookie", setCookie)

}
