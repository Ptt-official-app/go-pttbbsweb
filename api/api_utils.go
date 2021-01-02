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
	pttbbsapi "github.com/Ptt-official-app/go-pttbbs/api"
	"github.com/Ptt-official-app/go-pttbbs/bbs"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"gopkg.in/square/go-jose.v2"
	"gopkg.in/square/go-jose.v2/jwt"
)

func verifyJwt(c *gin.Context) (userID bbs.UUserID, err error) {
	jwt := pttbbsapi.GetJwt(c)

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

	if cl.Expire == nil {
		return false
	}

	currentNanoTS := jwt.NewNumericDate(time.Now())
	if *currentNanoTS > *cl.Expire {
		return false
	}

	return true
}

func processCSRFContent(filename string, c *gin.Context) {
	if !isValidOriginReferer(c) {
		processResult(c, nil, 403, ErrInvalidOrigin)
		return
	}

	file, err := os.Open(filename)
	if err != nil {
		logrus.Errorf("processCSRFContent: unable to open file: filename: %v", filename)
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
	logrus.Infof("processCSRFContent: filename: %v ext: %v mimeType: %v", filename, ext, mimeType)

	content := string(contentBytes)

	csrfToken, err := createCSRFToken()
	content = strings.Replace(content, "__CSRFTOKEN__", csrfToken, 1)

	c.Header("Cache-Control", "max-age="+strconv.Itoa(types.CSRF_TOKEN_TS))

	processStringResult(c, content, mimeType)
}
