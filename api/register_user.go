package api

import (
	pttbbsapi "github.com/Ptt-official-app/go-pttbbs/api"
	"github.com/Ptt-official-app/go-pttbbs/bbs"
	"github.com/Ptt-official-app/pttbbs-backend/schema"
	"github.com/Ptt-official-app/pttbbs-backend/types"
	"github.com/Ptt-official-app/pttbbs-backend/utils"
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

	Email    string `json:"email" form:"email"`
	Nickname string `json:"nickname,omitempty" form:"nickname,omitempty"`
	Realname string `json:"realname,omitempty" form:"realname,omitempty"`
	Career   string `json:"career,omitempty" form:"career,omitempty"`
}

func NewRegisterUserParams() *RegisterUserParams {
	return &RegisterUserParams{}
}

type RegisterUserResult LoginResult

func RegisterUserWrapper(c *gin.Context) {
	params := NewRegisterUserParams()
	FormJSON(RegisterUser, params, c)
}

func RegisterUser(remoteAddr string, user *UserInfo, params interface{}, c *gin.Context) (result interface{}, statusCode int, err error) {
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

	if types.IS_2FA {
		err = check2FAToken(bbs.UUserID(theParams.Username), theParams.Email, theParams.TwoFactorToken)
		if err != nil {
			return nil, 403, err
		}
	}

	// create db-record first to avoid race-condition
	updateNanoTS := types.NowNanoTS()
	userID := bbs.UUserID(theParams.Username)
	err = schema.CreateUserEmail(userID, theParams.Email, updateNanoTS)
	if err != nil {
		return nil, 403, err
	}

	// backend register
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

	// update db
	updateNanoTS = types.NowNanoTS()
	err = schema.UpdateUserEmailIsSet(userID, theParams.Email, true, updateNanoTS)
	if err != nil {
		return nil, statusCode, err
	}

	// result
	result = NewRegisterUserResult(result_b)

	setTokenToCookie(c, result_b.Jwt)

	return result, 200, nil
}

func NewRegisterUserResult(result_b *pttbbsapi.RegisterResult) *RegisterUserResult {
	return &RegisterUserResult{
		UserID:        result_b.UserID,
		AccessToken:   result_b.Jwt,
		TokenType:     "bearer",
		RefreshToken:  result_b.Refresh,
		AccessExpire:  types.Time8(result_b.AccessExpire),
		RefreshExpire: types.Time8(result_b.RefreshExpire),

		TokenUser: result_b.UserID,
	}
}
