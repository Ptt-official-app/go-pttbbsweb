package api

import (
	"github.com/Ptt-official-app/go-pttbbs/bbs"
	"github.com/Ptt-official-app/pttbbs-backend/types"
	"github.com/gin-gonic/gin"
)

type APIFunc func(remoteAddr string, user *UserInfo, params interface{}, c *gin.Context) (result interface{}, statusCode int, err error)

type PathAPIFunc func(remoteAddr string, user *UserInfo, params interface{}, path interface{}, c *gin.Context) (result interface{}, statusCode int, err error)

type LoginRequiredAPIFunc func(remoteAddr string, user *UserInfo, params interface{}, c *gin.Context) (result interface{}, statusCode int, err error)

type LoginRequiredPathAPIFunc func(remoteAddr string, user *UserInfo, params interface{}, path interface{}, c *gin.Context) (result interface{}, statusCode int, err error)

type RedirectPathAPIFunc func(remoteAddr string, params interface{}, path interface{}, c *gin.Context) (redirectPath string, statusCode int)

type errResult struct {
	Msg string

	TokenUser bbs.UUserID `json:"tokenuser"`
}

type ClientInfo struct {
	ClientID   string           `json:"c"`
	ClientType types.ClientType `json:"t"`
}

type UserInfo struct {
	UserID   bbs.UUserID
	IsOver18 bool
}
