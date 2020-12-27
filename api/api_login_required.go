package api

import (
	"strings"

	pttbbsapi "github.com/Ptt-official-app/go-pttbbs/api"

	"github.com/gin-gonic/gin"
)

func LoginRequiredJSON(theFunc LoginRequiredApiFunc, params interface{}, c *gin.Context) {
	err := c.ShouldBindJSON(params)
	if err != nil {
		processResult(c, nil, 400, err)
		return
	}

	loginRequiredProcess(theFunc, params, c)
}

func LoginRequiredQuery(theFunc LoginRequiredApiFunc, params interface{}, c *gin.Context) {
	err := c.ShouldBindQuery(params)
	if err != nil {
		processResult(c, nil, 400, err)
		return
	}

	loginRequiredProcess(theFunc, params, c)
}

func loginRequiredProcess(theFunc LoginRequiredApiFunc, params interface{}, c *gin.Context) {

	remoteAddr := strings.TrimSpace(c.ClientIP())
	if !isValidRemoteAddr(remoteAddr) {
		processResult(c, nil, 400, ErrInvalidRemoteAddr)
		return
	}

	tokenStr := strings.TrimSpace(c.GetHeader("Authorization"))
	tokenList := strings.Split(tokenStr, " ")
	if len(tokenList) != 2 {
		processResult(c, nil, 400, ErrInvalidToken)
		return
	}
	jwt := tokenList[1]

	userID, err := pttbbsapi.VerifyJwt(jwt)
	if err != nil {
		processResult(c, nil, 401, err)
		return
	}

	result, statusCode, err := theFunc(remoteAddr, userID, params, c)
	processResult(c, result, statusCode, err)
}
