package api

import (
	"strings"

	"github.com/Ptt-official-app/go-pttbbsweb/types"

	"github.com/Ptt-official-app/go-pttbbsweb/schema"

	pttbbsapi "github.com/Ptt-official-app/go-pttbbs/api"
	"github.com/Ptt-official-app/go-pttbbs/bbs"
	"github.com/gin-gonic/gin"
)

func LoginRequiredPathJSON(theFunc LoginRequiredPathAPIFunc, params interface{}, path interface{}, c *gin.Context) {
	err := c.ShouldBindJSON(params)
	if err != nil {
		processResult(c, nil, 400, err, "")
		return
	}

	loginRequiredPathProcess(theFunc, params, path, c)
}

func LoginRequiredPathQuery(theFunc LoginRequiredPathAPIFunc, params interface{}, path interface{}, c *gin.Context) {
	err := c.ShouldBindQuery(params)
	if err != nil {
		processResult(c, nil, 400, err, "")
		return
	}

	loginRequiredPathProcess(theFunc, params, path, c)
}

func loginRequiredPathProcess(theFunc LoginRequiredPathAPIFunc, params interface{}, path interface{}, c *gin.Context) {
	err := c.ShouldBindUri(path)
	if err != nil {
		processResult(c, nil, 400, err, "")
		return
	}

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
	userVisit := &schema.UserVisit{
		UserID:       userID,
		Action:       c.Request.Method + ":" + c.Request.URL.Path,
		UpdateNanoTS: types.NowNanoTS(),
	}
	_ = schema.UpdateUserVisit(userVisit)

	result, statusCode, err := theFunc(remoteAddr, userID, params, path, c)
	processResult(c, result, statusCode, err, userID)
}
