package api

import (
	"github.com/Ptt-official-app/go-openbbsmiddleware/schema"
	"github.com/Ptt-official-app/go-openbbsmiddleware/types"
	"github.com/Ptt-official-app/go-openbbsmiddleware/utils"
	pttbbsapi "github.com/Ptt-official-app/go-pttbbs/api"
	"github.com/Ptt-official-app/go-pttbbs/bbs"
	"github.com/gin-gonic/gin"
)

const REGISTER_USER_R = "/account/register"

type RegisterUserParams struct {
	ClientID     string `json:"client_id" form:"client_id"`
	ClientSecret string `json:"client_secret" form:"client_secret"`

	Username        string `json:"username" form:"username"`
	Password        string `json:"password" form:"password"`
	PasswordConfirm string `json:"password_confirm" form:"password_confirm"`

	TwoFactorToken string `json:"token" form:"token"`

	Over18 bool `json:"over18,omitempty" form:"over18,omitempty"`

	Email    string `json:"email,omitempty" form:"email,omitempty"`
	Nickname string `json:"nickname,omitempty" form:"nickname,omitempty"`
	Realname string `json:"realname,omitempty" form:"realname,omitempty"`
	Career   string `json:"career,omitempty" form:"career,omitempty"`
}

func NewRegisterUserParams() *RegisterUserParams {
	return &RegisterUserParams{}
}

type RegisterUserResult struct {
	UserID      bbs.UUserID `json:"user_id"`
	AccessToken string      `json:"access_token"`
	TokenType   string      `json:"token_type"`
}

func RegisterUserWrapper(c *gin.Context) {
	params := NewRegisterUserParams()
	FormJSON(RegisterUser, params, c)
}

func RegisterUser(remoteAddr string, params interface{}, c *gin.Context) (result interface{}, statusCode int, err error) {
	theParams, ok := params.(*RegisterUserParams)
	if !ok {
		return nil, 400, ErrInvalidParams
	}

	if theParams.Password != theParams.PasswordConfirm {
		return nil, 400, ErrInvalidParams
	}

	isValidClient, client := checkClient(theParams.ClientID, theParams.ClientSecret)

	if !isValidClient {
		return nil, 400, ErrInvalidParams
	}

	clientInfo := getClientInfo(client)

	err = check2FAToken(bbs.UUserID(theParams.Username), theParams.TwoFactorToken)
	if err != nil {
		return nil, 403, err
	}

	//backend register
	theParams_b := &pttbbsapi.RegisterParams{
		ClientInfo: clientInfo,
		Username:   theParams.Username,
		Passwd:     theParams.Password,
		Over18:     theParams.Over18,

		Email:    theParams.Email,
		Nickname: types.Utf8ToBig5(theParams.Nickname),
		Realname: types.Utf8ToBig5(theParams.Realname),
		Career:   types.Utf8ToBig5(theParams.Career),
	}
	var result_b *pttbbsapi.RegisterResult

	url := pttbbsapi.REGISTER_R
	statusCode, err = utils.BackendPost(c, url, theParams_b, nil, &result_b)
	if err != nil || statusCode != 200 {
		return nil, statusCode, err
	}

	//update db
	updateNanoTS := types.NowNanoTS()
	accessToken_db, err := deserializeAccessTokenAndUpdateDB(result_b.UserID, result_b.Jwt, updateNanoTS)
	if err != nil {
		return nil, 500, err
	}

	//result
	result = NewRegisterUserResult(accessToken_db)

	setTokenToCookie(c, accessToken_db.AccessToken)

	return result, 200, nil
}

func NewRegisterUserResult(accessToken_db *schema.AccessToken) *RegisterUserResult {
	return &RegisterUserResult{
		UserID:      accessToken_db.UserID,
		AccessToken: accessToken_db.AccessToken,
		TokenType:   "bearer",
	}
}
