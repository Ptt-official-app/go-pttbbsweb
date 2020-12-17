package main

import (
	"strings"

	"github.com/Ptt-official-app/go-openbbsmiddleware/api"

	"github.com/gin-gonic/gin"
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

func (a *Api) Query(c *gin.Context) {
	err := c.ShouldBindQuery(a.Params)
	if err != nil {
		processResult(c, nil, 400, err)
		return
	}

	a.process(c)

}

func (a *Api) Json(c *gin.Context) {
	err := c.ShouldBindJSON(a.Params)
	if err != nil {
		processResult(c, nil, 400, err)
		return
	}

	a.process(c)

}

func (a *Api) process(c *gin.Context) {
	remoteAddr := c.ClientIP()
	if !isValidRemoteAddr(remoteAddr) {
		processResult(c, nil, 400, ErrInvalidRemoteAddr)
		return
	}

	result, statusCode, err := a.Func(remoteAddr, a.Params, c)
	processResult(c, result, statusCode, err)
}

func NewLoginRequiredApi(f api.LoginRequiredApiFunc, params interface{}) *LoginRequiredApi {
	return &LoginRequiredApi{Func: f, Params: params}
}

func (a *LoginRequiredApi) Query(c *gin.Context) {
	err := c.ShouldBindQuery(a.Params)
	if err != nil {
		processResult(c, nil, 400, err)
		return
	}

	a.process(c)
}

func (a *LoginRequiredApi) Json(c *gin.Context) {
	err := c.ShouldBindJSON(a.Params)
	if err != nil {
		processResult(c, nil, 400, err)
		return
	}

	a.process(c)
}

func (a *LoginRequiredApi) process(c *gin.Context) {
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

	userID, err := api.VerifyJwt(jwt)
	if err != nil {
		processResult(c, nil, 401, err)
		return
	}

	result, statusCode, err := a.Func(remoteAddr, userID, a.Params, c)
	processResult(c, result, statusCode, err)
}

func NewLoginRequiredPathApi(f api.LoginRequiredPathApiFunc, params interface{}, path interface{}) *LoginRequiredPathApi {
	return &LoginRequiredPathApi{Func: f, Params: params, Path: path}
}

func (a *LoginRequiredPathApi) Query(c *gin.Context) {
	err := c.ShouldBindQuery(a.Params)
	if err != nil {
		processResult(c, nil, 400, err)
		return
	}
	err = c.ShouldBindUri(a.Path)
	if err != nil {
		processResult(c, nil, 400, err)
	}

	a.process(c)
}

func (a *LoginRequiredPathApi) Json(c *gin.Context) {
	err := c.ShouldBindJSON(a.Params)
	if err != nil {
		processResult(c, nil, 400, err)
		return
	}

	err = c.ShouldBindUri(a.Path)
	if err != nil {
		processResult(c, nil, 400, err)
	}

	a.process(c)
}

func (a *LoginRequiredPathApi) process(c *gin.Context) {
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

	userID, err := api.VerifyJwt(jwt)
	if err != nil {
		processResult(c, nil, 401, err)
		return
	}

	result, statusCode, err := a.Func(remoteAddr, userID, a.Params, a.Path, c)
	processResult(c, result, statusCode, err)
}

func processResult(c *gin.Context, result interface{}, statusCode int, err error) {
	if statusCode == 200 {
		c.JSON(200, result)
	} else {
		c.JSON(statusCode, &errResult{err.Error()})
	}
}
