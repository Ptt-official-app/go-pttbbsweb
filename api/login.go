package api

import (
	"fmt"

	"github.com/Ptt-official-app/go-openbbsmiddleware/schema"
	"github.com/Ptt-official-app/go-openbbsmiddleware/types"
	"github.com/Ptt-official-app/go-openbbsmiddleware/utils"
	pttbbsapi "github.com/Ptt-official-app/go-pttbbs/api"
	"github.com/Ptt-official-app/go-pttbbs/bbs"
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
	UserID      bbs.UUserID `json:"user_id"`
	AccessToken string      `json:"access_token"`
	TokenType   string      `json:"token_type"`
}

// LoginLog record user login info, no matter success or not
type LoginLog struct {
	ClientInfo
	LoginId   string
	LoginTime types.NanoTS
	LoginIP   string
	IsSuccess bool
}

func (l LoginLog) Stringer() string {
	var success string
	if l.IsSuccess {
		success = "\033[97;42mSuccess\033[0m"
	} else {
		success = "\033[97;41mFail\033[0m"
	}
	return fmt.Sprintf("ID: %s login %s from %s at %v Client: %v \n", l.LoginId, success, l.LoginIP, l.LoginTime.ToTime(), l.ClientInfo)
}

func LoginWrapper(c *gin.Context) {
	params := NewLoginParams()
	FormJSON(Login, params, c)
}

func Login(remoteAddr string, params interface{}, c *gin.Context) (result interface{}, statusCode int, err error) {
	theParams, ok := params.(*LoginParams)
	// record user login
	loginLog := LoginLog{
		ClientInfo: ClientInfo{
			ClientID: theParams.ClientID,
		},
		LoginId:   theParams.Username,
		LoginIP:   remoteAddr,
		LoginTime: types.NowNanoTS(),
		IsSuccess: false, // default is false
	}
	defer func() {
		logrus.Info(loginLog.Stringer())
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

	// update db
	updateNanoTS := types.NowNanoTS()
	accessToken_db, err := deserializeAccessTokenAndUpdateDB(result_b.UserID, result_b.Jwt, updateNanoTS)
	if err != nil {
		return nil, 500, err
	}

	// result
	result = NewLoginResult(accessToken_db)

	setTokenToCookie(c, accessToken_db.AccessToken)

	return result, 200, nil
}

func NewLoginResult(accessToken_db *schema.AccessToken) *LoginResult {
	return &LoginResult{
		UserID:      accessToken_db.UserID,
		AccessToken: accessToken_db.AccessToken,
		TokenType:   "bearer",
	}
}
