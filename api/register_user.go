package api

import (
	"github.com/Ptt-official-app/go-openbbsmiddleware/backend"
	"github.com/Ptt-official-app/go-openbbsmiddleware/schema"
	"github.com/Ptt-official-app/go-openbbsmiddleware/utils"
	"github.com/gin-gonic/gin"
)

const REGISTER_USER_R = "/account/register"

type RegisterUserParams struct {
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`

	Username        string `json:"username"`
	Password        string `json:"password"`
	PasswordConfirm string `json:"password_confirm"`

	Over18 bool `json:"over18,omitempty"`

	Email    string `json:"email,omitempty"`
	Nickname string `json:"nickname,omitempty"`
	Realname string `json:"realname,omitempty"`
	Career   string `json:"career,omitempty"`
	Address  string `json:"address,omitempty"`
}

func NewRegisterUserParams() *RegisterUserParams {
	return &RegisterUserParams{}
}

type RegisterUserResult struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
}

func RegisterUser(remoteAddr string, params interface{}, c *gin.Context) (result interface{}, statusCode int, err error) {
	theParams, ok := params.(*RegisterUserParams)
	if !ok {
		return nil, 400, ErrInvalidParams
	}

	if theParams.Password != theParams.PasswordConfirm {
		return nil, 400, ErrInvalidParams
	}

	if !isValidClient(theParams.ClientID, theParams.ClientSecret) {
		return nil, 400, ErrInvalidParams
	}

	//backend register
	theParams_b := &backend.RegisterParams{
		UserID:   theParams.Username,
		Passwd:   theParams.Password,
		Over18:   theParams.Over18,
		Email:    theParams.Email,
		Nickname: theParams.Nickname,
	}
	var result_b *backend.RegisterResult

	url := backend.WithPrefix(backend.REGISTER_R)
	statusCode, err = utils.HttpPost(c, url, theParams_b, nil, &result_b)
	if err != nil || statusCode != 200 {
		return nil, statusCode, err
	}

	//update db
	userID, err := VerifyJwt(result_b.Jwt)
	if err != nil {
		return nil, 401, err
	}

	nowNanoTS := utils.GetNowNanoTS()
	query := &schema.AccessToken{
		AccessToken:  result_b.Jwt,
		UserID:       userID,
		UpdateNanoTS: nowNanoTS,
	}

	//TODO: possibly change to insert？～
	_, err = schema.AccessToken_c.Update(query, query)
	if err != nil {
		return nil, 500, err
	}

	//result
	result = &RegisterUserResult{
		AccessToken: result_b.Jwt,
		TokenType:   result_b.TokenType,
	}

	return result, 200, nil
}
