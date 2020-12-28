package api

import (
	"strings"

	pttbbsapi "github.com/Ptt-official-app/go-pttbbs/api"
	"github.com/gin-gonic/gin"
)

func LoginRequiredPathJSON(theFunc LoginRequiredPathApiFunc, params interface{}, path interface{}, c *gin.Context) {
	err := c.ShouldBindJSON(params)
	if err != nil {
		processResult(c, nil, 400, err)
		return
	}

	loginRequiredPathProcess(theFunc, params, path, c)
}

func LoginRequiredPathQuery(theFunc LoginRequiredPathApiFunc, params interface{}, path interface{}, c *gin.Context) {
	err := c.ShouldBindQuery(params)
	if err != nil {
		processResult(c, nil, 400, err)
		return
	}

	loginRequiredPathProcess(theFunc, params, path, c)
}

func loginRequiredPathProcess(theFunc LoginRequiredPathApiFunc, params interface{}, path interface{}, c *gin.Context) {

	err := c.ShouldBindUri(path)
	if err != nil {
		processResult(c, nil, 400, err)
		return
	}

	remoteAddr := strings.TrimSpace(c.ClientIP())
	if !isValidRemoteAddr(remoteAddr) {
		processResult(c, nil, 400, ErrInvalidRemoteAddr)
		return
	}

	jwt := pttbbsapi.GetJwt(c)

	userID, err := pttbbsapi.VerifyJwt(jwt)
	if err != nil {
		processResult(c, nil, 401, err)
		return
	}

	result, statusCode, err := theFunc(remoteAddr, userID, params, path, c)
	processResult(c, result, statusCode, err)

}
