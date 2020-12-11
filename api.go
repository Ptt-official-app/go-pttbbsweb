package main

import (
	"strings"
	"time"

	"github.com/Ptt-official-app/go-openbbsmiddleware/api"

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

func NewApi(f api.ApiFunc, params interface{}) *Api {
	return &Api{Func: f, Params: params}
}

func NewLoginRequiredApi(f api.LoginRequiredApiFunc, params interface{}) *LoginRequiredApi {
	return &LoginRequiredApi{Func: f, Params: params}
}

func (api *Api) Query(c *gin.Context) {
	err := c.ShouldBindQuery(api.Params)
	if err != nil {
		processResult(c, nil, err)
		return
	}

	api.process(c)

}

func (api *Api) Json(c *gin.Context) {
	err := c.ShouldBindJSON(api.Params)
	if err != nil {
		processResult(c, nil, err)
		return
	}

	api.process(c)

}

func (api *Api) process(c *gin.Context) {
	remoteAddr := c.ClientIP()
	if !isValidRemoteAddr(remoteAddr) {
		processResult(c, nil, ErrInvalidRemoteAddr)
		return
	}

	result, err := api.Func(remoteAddr, api.Params)
	processResult(c, result, err)
}

func (api *LoginRequiredApi) LoginRequiredQuery(c *gin.Context) {
	err := c.ShouldBindQuery(api.Params)
	if err != nil {
		processResult(c, nil, err)
		return
	}

	api.process(c)
}

func (api *LoginRequiredApi) LoginRequiredJson(c *gin.Context) {
	err := c.ShouldBindJSON(api.Params)
	if err != nil {
		processResult(c, nil, err)
		return
	}

	api.process(c)
}

func (api *LoginRequiredApi) process(c *gin.Context) {
	//https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/X-Forwarded-For
	remoteAddr := c.ClientIP()
	if !isValidRemoteAddr(remoteAddr) {
		processResult(c, nil, ErrInvalidRemoteAddr)
		return
	}

	tokenStr := strings.TrimSpace(c.GetHeader("Authorization"))
	tokenList := strings.Split(tokenStr, " ")
	if len(tokenList) != 2 {
		processResult(c, nil, ErrInvalidToken)
		return
	}
	jwt := tokenList[1]

	userID, err := verifyJwt(jwt)
	if err != nil {
		processResult(c, nil, err)
		return
	}

	result, err := api.Func(remoteAddr, userID, api.Params)
	processResult(c, result, err)
}

//verifyJwt
//
//from https://github.com/Ptt-official-app/go-pttbbs/blob/main/api.go#L93
func verifyJwt(raw string) (userID string, err error) {
	tok, err := jwt.ParseSigned(raw)
	if err != nil {
		return "", ErrInvalidToken
	}

	cl := &api.JwtClaim{}
	if err := tok.Claims(api.JWT_SECRET, cl); err != nil {
		return "", ErrInvalidToken
	}

	currentNanoTS := jwt.NewNumericDate(time.Now())
	if *currentNanoTS > *cl.Expire {
		return "", ErrInvalidToken
	}

	return cl.UserID, nil
}

func processResult(c *gin.Context, result interface{}, err error) {
	switch err {
	case nil:
		c.JSON(200, result)

	case ErrInvalidHost:
		c.JSON(400, &errResult{err.Error()})
	case ErrInvalidRemoteAddr:
		c.JSON(400, &errResult{err.Error()})
	case api.ErrInvalidParams:
		c.JSON(400, &errResult{err.Error()})

	case ErrInvalidToken:
		c.JSON(401, &errResult{err.Error()})
	case api.ErrLoginFailed:
		c.JSON(401, &errResult{err.Error()})
	default:
		c.JSON(500, &errResult{err.Error()})
	}
}
