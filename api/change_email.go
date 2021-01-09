package api

import (
	pttbbsapi "github.com/Ptt-official-app/go-pttbbs/api"

	"github.com/Ptt-official-app/go-openbbsmiddleware/utils"
	"github.com/Ptt-official-app/go-pttbbs/bbs"
	"github.com/gin-gonic/gin"
)

const CHANGE_EMAIL_R = "/user/:user_id/changeemail"

type ChangeEmailParams struct {
	ClientID     string `json:"client_id" form:"client_id"`
	ClientSecret string `json:"client_secret" form:"client_secret"`
	Jwt          string `json:"token" form:"token" url:"token"`
}

type ChangeEmailPath struct {
	UserID bbs.UUserID `uri:"user_id"`
}

type ChangeEmailResult struct {
	Email string `json:"email"`
}

func ChangeEmailWrapper(c *gin.Context) {
	params := &ChangeEmailParams{}
	path := &ChangeEmailPath{}

	LoginRequiredPathJSON(ChangeEmail, params, path, c)
}

func ChangeEmail(remoteAddr string, userID bbs.UUserID, params interface{}, path interface{}, c *gin.Context) (result interface{}, statusCode int, err error) {
	theParams, ok := params.(*ChangeEmailParams)
	if !ok {
		return nil, 400, ErrInvalidParams
	}

	thePath, ok := path.(*ChangeEmailPath)
	if !ok {
		return nil, 400, ErrInvalidPath
	}

	//get backend data
	theParams_b := &pttbbsapi.ChangeEmailParams{
		Jwt: theParams.Jwt,
	}

	var result_b *pttbbsapi.ChangeEmailResult

	urlMap := map[string]string{
		"uid": string(thePath.UserID),
	}
	url := utils.MergeURL(urlMap, pttbbsapi.CHANGE_EMAIL_R)

	statusCode, err = utils.BackendPost(c, url, theParams_b, nil, &result_b)
	if err != nil {
		return nil, statusCode, err
	}

	return &ChangeEmailResult{Email: result_b.Email}, 200, nil

}

func changeEmailRedirectURL(userID bbs.UUserID) (url string) {
	urlMap := map[string]string{
		"user_id": string(userID),
	}
	return utils.MergeURL(urlMap, GET_USER_INFO_R)
}
