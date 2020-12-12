package api

import (
	"github.com/Ptt-official-app/go-openbbsmiddleware/backend"
	"github.com/Ptt-official-app/go-openbbsmiddleware/schema"
	"github.com/Ptt-official-app/go-openbbsmiddleware/utils"
	"github.com/gin-gonic/gin"
)

type LoginParams struct {
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
	Username     string `json:"username"`
	Password     string `json:"password"`
}

type LoginResult struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
}

func Login(remoteAddr string, params interface{}, c *gin.Context) (result interface{}, statusCode int, err error) {
	theParams, ok := params.(*LoginParams)
	if !ok {
		return nil, 400, ErrInvalidParams
	}

	if !isValidClient(theParams.ClientID, theParams.ClientSecret) {
		return nil, 400, ErrInvalidParams
	}

	//backend login
	loginParams_b := &backend.LoginParams{Username: theParams.Username, Password: theParams.Password}
	result_b := &backend.LoginResults{}

	url := backend.WithPrefix(backend.LOGIN_R)
	statusCode, err = utils.HttpPost(c, url, loginParams_b, nil, result_b)
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
	result = &LoginResult{
		AccessToken: result_b.AccessToken,
		TokenType:   result_b.TokenType,
	}

	return result, 200, nil
}
