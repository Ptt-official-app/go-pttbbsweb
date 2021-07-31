package api

import (
	"github.com/Ptt-official-app/go-openbbsmiddleware/types"
	"github.com/Ptt-official-app/go-pttbbs/bbs"
	"github.com/gin-gonic/gin"
)

type APIFunc func(remoteAddr string, params interface{}, c *gin.Context) (result interface{}, statusCode int, err error)

type LoginRequiredAPIFunc func(remoteAddr string, userID bbs.UUserID, params interface{}, c *gin.Context) (result interface{}, statusCode int, err error)

type LoginRequiredPathAPIFunc func(remoteAddr string, userID bbs.UUserID, params interface{}, path interface{}, c *gin.Context) (result interface{}, statusCode int, err error)

type RedirectPathAPIFunc func(remoteAddr string, params interface{}, path interface{}, c *gin.Context) (redirectPath string, statusCode int)

type errResult struct {
	Msg string
}

type ClientInfo struct {
	ClientID   string           `json:"c"`
	ClientType types.ClientType `json:"t"`
}
