package api

import (
	"github.com/Ptt-official-app/go-openbbsmiddleware/types"
	"github.com/Ptt-official-app/go-openbbsmiddleware/utils"
	pttbbsapi "github.com/Ptt-official-app/go-pttbbs/api"
	"github.com/Ptt-official-app/go-pttbbs/bbs"
	"github.com/gin-gonic/gin"
)

const ATTEMPT_REGISTER_USER_R = "/account/attemptregister"

type AttemptRegisterUserParams struct {
	ClientID     string `json:"client_id" form:"client_id"`
	ClientSecret string `json:"client_secret" form:"client_secret"`

	Username string `json:"username" form:"username"`
	Email    string `json:"email" form:"email"`
}

type AttemptRegisterUserResult struct {
	Username string `json:"user_id"`
}

func AttemptRegisterUserWrapper(c *gin.Context) {
	params := &AttemptRegisterUserParams{}
	FormJSON(AttemptRegisterUser, params, c)
}

func AttemptRegisterUser(remoteAddr string, params interface{}, c *gin.Context) (result interface{}, statusCode int, err error) {

	theParams, ok := params.(*AttemptRegisterUserParams)
	if !ok {
		return nil, 400, ErrInvalidParams
	}

	//check existing user
	theParams_b := &pttbbsapi.CheckExistsUserParams{
		Username: theParams.Username,
	}
	var result_b *pttbbsapi.CheckExistsUserResult

	url := pttbbsapi.CHECK_EXISTS_USER_R
	statusCode, err = utils.BackendPost(c, url, theParams_b, nil, &result_b)
	if err != nil || statusCode != 200 {
		return nil, statusCode, err
	}
	if result_b.IsExists {
		return nil, 400, ErrAlreadyExists
	}

	err = gen2FATokenAndSendEmail(bbs.UUserID(theParams.Username), theParams.Email, types.ATTEMPT_REGISTER_USER_TITLE, types.ATTEMPT_REGISTER_USER_TEMPLATE_CONTENT)
	if err != nil {
		return nil, 400, err
	}

	return &AttemptRegisterUserResult{Username: theParams.Username}, 200, nil
}
