package api

import (
	"strings"

	pttbbsapi "github.com/Ptt-official-app/go-pttbbs/api"
	"github.com/Ptt-official-app/go-pttbbs/bbs"
	"github.com/gin-gonic/gin"
)

func LoginRequiredJSON(theFunc LoginRequiredAPIFunc, params interface{}, c *gin.Context) {
	err := c.ShouldBindJSON(params)
	if err != nil {
		processResult(c, nil, 400, err, "")
		return
	}

	loginRequiredProcess(theFunc, params, c)
}

func LoginRequiredQuery(theFunc LoginRequiredAPIFunc, params interface{}, c *gin.Context) {
	err := c.ShouldBindQuery(params)
	if err != nil {
		processResult(c, nil, 400, err, "")
		return
	}

	loginRequiredProcess(theFunc, params, c)
}

func loginRequiredProcess(theFunc LoginRequiredAPIFunc, params interface{}, c *gin.Context) {
	remoteAddr := strings.TrimSpace(c.ClientIP())
	if !isValidRemoteAddr(remoteAddr) {
		processResult(c, nil, 403, ErrInvalidRemoteAddr, "")
		return
	}

	if !isValidOriginReferer(c) {
		processResult(c, nil, 403, ErrInvalidOrigin, "")
		return
	}

	userID, err := verifyJwt(c)
	if err != nil {
		userID = bbs.UUserID(pttbbsapi.GUEST)
	}

	isOver18 := verifyIsOver18(c)

	user := &UserInfo{IsOver18: isOver18, UserID: userID}

	result, statusCode, err := theFunc(remoteAddr, user, params, c)
	processResult(c, result, statusCode, err, userID)
}
