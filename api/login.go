package api

import (
	"fmt"

	pttbbsapi "github.com/Ptt-official-app/go-pttbbs/api"
	"github.com/Ptt-official-app/go-pttbbs/bbs"
	"github.com/Ptt-official-app/pttbbs-backend/types"
	"github.com/Ptt-official-app/pttbbs-backend/utils"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

const LOGIN_R = "/account/login"

type LoginParams struct {
	ClientID     string `json:"client_id" form:"client_id"`
	ClientSecret string `json:"client_secret" form:"client_secret"`

	Username string `json:"username" form:"username"`
	Password string `json:"password" form:"password"`
}

func NewLoginParams() *LoginParams {
	return &LoginParams{}
}

type LoginResult struct {
	UserID        bbs.UUserID `json:"user_id"`
	AccessToken   string      `json:"access_token"`
	TokenType     string      `json:"token_type"`
	RefreshToken  string      `json:"refresh_token"`
	AccessExpire  types.Time8 `json:"access_expire"`
	RefreshExpire types.Time8 `json:"refresh_expire"`

	TokenUser bbs.UUserID `json:"tokenuser"`
}

// LoginLog record user login info, no matter success or not
type LoginLog struct {
	ClientInfo
	LoginID   string
	LoginTime types.NanoTS
	LoginIP   string
	IsSuccess bool
}

func (l *LoginLog) String() string {
	var success string
	if l.IsSuccess {
		success = "\033[97;42mSuccess\033[0m"
	} else {
		success = "\033[97;41mFail\033[0m"
	}
	return fmt.Sprintf("ID: %s login %s from %s at %v Client: %v \n", l.LoginID, success, l.LoginIP, l.LoginTime.ToTime(), l.ClientInfo)
}

func LoginWrapper(c *gin.Context) {
	params := NewLoginParams()
	FormJSON(Login, params, c)
}

func Login(remoteAddr string, params interface{}, c *gin.Context) (result interface{}, statusCode int, err error) {
	theParams, ok := params.(*LoginParams)
	// record user login
	loginLog := &LoginLog{
		ClientInfo: ClientInfo{
			ClientID: theParams.ClientID,
		},
		LoginID:   theParams.Username,
		LoginIP:   remoteAddr,
		LoginTime: types.NowNanoTS(),
		IsSuccess: false, // default is false
	}
	defer func() {
		logrus.Infof("%v", loginLog)
	}()

	if !ok {
		return nil, 400, ErrInvalidParams
	}

	isValidClient, client := checkClient(theParams.ClientID, theParams.ClientSecret)

	if !isValidClient {
		return nil, 400, ErrInvalidParams
	}

	clientInfo := getClientInfo(client)

	// backend login
	theParams_b := &pttbbsapi.LoginParams{
		ClientInfo: clientInfo,
		Username:   theParams.Username,
		Passwd:     theParams.Password,
	}

	var result_b *pttbbsapi.LoginResult

	url := pttbbsapi.LOGIN_R
	statusCode, err = utils.BackendPost(c, url, theParams_b, nil, &result_b)

	if err != nil || statusCode != 200 {
		return nil, statusCode, err
	}
	// update: loginLog success login
	loginLog.IsSuccess = true

	// result
	result = NewLoginResult(result_b)

	setTokenToCookie(c, result_b.Jwt)

	return result, 200, nil
}

func NewLoginResult(result_b *pttbbsapi.LoginResult) *LoginResult {
	return &LoginResult{
		UserID:        result_b.UserID,
		AccessToken:   result_b.Jwt,
		TokenType:     "bearer",
		RefreshToken:  result_b.Refresh,
		AccessExpire:  types.Time8(result_b.AccessExpire),
		RefreshExpire: types.Time8(result_b.RefreshExpire),
		TokenUser:     result_b.UserID,
	}
}
