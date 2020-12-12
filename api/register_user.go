package api

import (
	"github.com/Ptt-official-app/go-openbbsmiddleware/backend"
	"github.com/Ptt-official-app/go-openbbsmiddleware/schema"
	"github.com/Ptt-official-app/go-openbbsmiddleware/utils"
	"github.com/gin-gonic/gin"
)

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
	registerParams_b := &backend.RegisterParams{
		Username: theParams.Username,
		Password: theParams.Password,
		Over18:   theParams.Over18,
		Email:    theParams.Email,
		Nickname: theParams.Nickname,
	}
	result_b := &backend.RegisterResults{}

	url := backend.WithPrefix(backend.REGISTER_R)
	statusCode, err = utils.HttpPost(c, url, registerParams_b, nil, result_b)
	if err != nil || statusCode != 200 {
		return nil, statusCode, err
	}

	//update db
	nowNanoTS := utils.GetNowNanoTS()
	query := &schema.AccessToken{
		AccessToken:  result_b.AccessToken,
		UserID:       theParams.Username,
		UpdateNanoTS: nowNanoTS,
	}

	//TODO: possibly change to insert？～
	_, err = schema.AccessToken_c.Update(query, query)
	if err != nil {
		return nil, 500, err
	}

	//result
	result = &RegisterUserResult{
		AccessToken: result_b.AccessToken,
		TokenType:   result_b.TokenType,
	}

	return result, 200, nil
}
