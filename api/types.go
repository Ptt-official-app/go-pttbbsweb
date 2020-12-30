package api

import (
	"github.com/Ptt-official-app/go-pttbbs/bbs"
	"github.com/gin-gonic/gin"
)

type ApiFunc func(remoteAddr string, params interface{}, c *gin.Context) (result interface{}, statusCode int, err error)

type LoginRequiredApiFunc func(remoteAddr string, userID bbs.UUserID, params interface{}, c *gin.Context) (result interface{}, statusCode int, err error)

type LoginRequiredPathApiFunc func(remoteAddr string, userID bbs.UUserID, params interface{}, path interface{}, c *gin.Context) (result interface{}, statusCode int, err error)

type errResult struct {
	Msg string
}

type ClientInfo struct {
	ClientID   string     `json:"c"`
	ClientType ClientType `json:"t"`
}

type ClientType int

const (
	CLIENT_TYPE_APP ClientType = iota
	CLIENT_TYPE_WEB
)
