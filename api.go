package main

import (
	"strings"
	"time"

	"github.com/Ptt-official-app/go-openbbsmiddleware/api"
	"github.com/Ptt-official-app/go-openbbsmiddleware/types"

	"github.com/gin-gonic/gin"
	"gopkg.in/square/go-jose.v2/jwt"
)

type Api struct {
	Func   api.ApiFunc
	Params interface{}
}

type LoginRequiredApi struct {
	Func   api.LoginRequiredApiFunc
	Params interface{}
}

type LoginRequiredPathApi struct {
	Func   api.LoginRequiredPathApiFunc
	Params interface{}
	Path   interface{}
}

func NewApi(f api.ApiFunc, params interface{}) *Api {
	return &Api{Func: f, Params: params}
}

func (api *Api) Query(c *gin.Context) {
	err := c.ShouldBindQuery(api.Params)
	if err != nil {
		processResult(c, nil, 400, err)
		return
	}

	api.process(c)

}

func (api *Api) Json(c *gin.Context) {
	err := c.ShouldBindJSON(api.Params)
	if err != nil {
		processResult(c, nil, 400, err)
		return
	}

	api.process(c)

}

func (api *Api) process(c *gin.Context) {
	remoteAddr := c.ClientIP()
	if !isValidRemoteAddr(remoteAddr) {
		processResult(c, nil, 400, ErrInvalidRemoteAddr)
		return
	}

	result, statusCode, err := api.Func(remoteAddr, api.Params, c)
	processResult(c, result, statusCode, err)
}

func NewLoginRequiredApi(f api.LoginRequiredApiFunc, params interface{}) *LoginRequiredApi {
	return &LoginRequiredApi{Func: f, Params: params}
}

func (api *LoginRequiredApi) Query(c *gin.Context) {
	err := c.ShouldBindQuery(api.Params)
	if err != nil {
		processResult(c, nil, 400, err)
		return
	}

	api.process(c)
}

func (api *LoginRequiredApi) Json(c *gin.Context) {
	err := c.ShouldBindJSON(api.Params)
	if err != nil {
		processResult(c, nil, 400, err)
		return
	}

	api.process(c)
}

func (api *LoginRequiredApi) process(c *gin.Context) {
	//https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/X-Forwarded-For
	remoteAddr := c.ClientIP()
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

	userID, err := verifyJwt(jwt)
	if err != nil {
		processResult(c, nil, 401, err)
		return
	}

	result, statusCode, err := api.Func(remoteAddr, userID, api.Params, c)
	processResult(c, result, statusCode, err)
}

func NewLoginRequiredPathApi(f api.LoginRequiredPathApiFunc, params interface{}, path interface{}) *LoginRequiredPathApi {
	return &LoginRequiredPathApi{Func: f, Params: params, Path: path}
}

func (api *LoginRequiredPathApi) Query(c *gin.Context) {
	err := c.ShouldBindQuery(api.Params)
	if err != nil {
		processResult(c, nil, 400, err)
		return
	}
	err = c.ShouldBindUri(api.Path)
	if err != nil {
		processResult(c, nil, 400, err)
	}

	api.process(c)
}

func (api *LoginRequiredPathApi) Json(c *gin.Context) {
	err := c.ShouldBindJSON(api.Params)
	if err != nil {
		processResult(c, nil, 400, err)
		return
	}

	err = c.ShouldBindUri(api.Path)
	if err != nil {
		processResult(c, nil, 400, err)
	}

	api.process(c)
}

func (api *LoginRequiredPathApi) process(c *gin.Context) {
	//https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/X-Forwarded-For
	remoteAddr := c.ClientIP()
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

	userID, err := verifyJwt(jwt)
	if err != nil {
		processResult(c, nil, 401, err)
		return
	}

	result, statusCode, err := api.Func(remoteAddr, userID, api.Params, api.Path, c)
	processResult(c, result, statusCode, err)
}

//verifyJwt
//
//from https://github.com/Ptt-official-app/go-pttbbs/blob/main/api.go#L93
func verifyJwt(raw string) (userID string, err error) {
	tok, err := jwt.ParseSigned(raw)
	if err != nil {
		return "", ErrInvalidToken
	}

	cl := &types.JwtClaim{}
	if err := tok.Claims(types.JWT_SECRET, cl); err != nil {
		return "", ErrInvalidToken
	}

	currentNanoTS := jwt.NewNumericDate(time.Now())
	if *currentNanoTS > *cl.Expire {
		return "", ErrInvalidToken
	}

	return cl.UserID, nil
}

func processResult(c *gin.Context, result interface{}, statusCode int, err error) {
	if statusCode == 200 {
		c.JSON(200, result)
	} else {
		c.JSON(statusCode, &errResult{err.Error()})
	}
}
