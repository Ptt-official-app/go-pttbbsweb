package api

import (
	"github.com/Ptt-official-app/go-openbbsmiddleware/types"
	"github.com/Ptt-official-app/go-openbbsmiddleware/utils"
	pttbbsapi "github.com/Ptt-official-app/go-pttbbs/api"
	"github.com/Ptt-official-app/go-pttbbs/bbs"
	"github.com/gin-gonic/gin"
)

const ATTEMPT_SET_ID_EMAIL_R = "/user/:user_id/attemptsetidemail"

type AttemptSetIDEmailParams struct {
	ClientID     string `json:"client_id" form:"client_id"`
	ClientSecret string `json:"client_secret" form:"client_secret"`

	Password string `json:"password" form:"password"`
	Email    string `json:"email"`
}

type AttemptSetIDEmailPath struct {
	UserID bbs.UUserID `uri:"user_id"`
}

type AttemptSetIDEmailResult struct {
	UserID bbs.UUserID `json:"user_id"`
	Email  string      `json:"email"`
}

func AttemptSetIDEmailWrapper(c *gin.Context) {
	params := &AttemptSetIDEmailParams{}
	path := &AttemptSetIDEmailPath{}

	LoginRequiredPathJSON(AttemptSetIDEmail, params, path, c)
}

func AttemptSetIDEmail(remoteAddr string, userID bbs.UUserID, params interface{}, path interface{}, c *gin.Context) (result interface{}, statusCode int, err error) {

	theParams, ok := params.(*AttemptSetIDEmailParams)
	if !ok {
		return nil, 400, ErrInvalidParams
	}

	thePath, ok := path.(*AttemptSetIDEmailPath)
	if !ok {
		return nil, 400, ErrInvalidPath
	}

	err = checkUniqueIDEmail(theParams.Email)
	if err != nil {
		return nil, 403, err
	}

	isValidClient, client := checkClient(theParams.ClientID, theParams.ClientSecret)

	if !isValidClient {
		return nil, 400, ErrInvalidParams
	}

	clientInfo := getClientInfo(client)

	//get backend data
	theParams_b := &pttbbsapi.AttemptSetIDEmailParams{
		ClientInfo: clientInfo,
		Passwd:     theParams.Password,
		Email:      theParams.Email,
	}

	var result_b *pttbbsapi.AttemptSetIDEmailResult

	urlMap := map[string]string{
		"uid": string(thePath.UserID),
	}
	url := utils.MergeURL(urlMap, pttbbsapi.ATTEMPT_SET_ID_EMAIL_R)

	statusCode, err = utils.BackendPost(c, url, theParams_b, nil, &result_b)
	if err != nil {
		return nil, statusCode, err
	}

	err = deserializeEmailTokenAndEmail(theParams.Email, result_b.UserID, result_b.Jwt, USER_SET_ID_EMAIL_HTML_R, types.IDEMAILTOKEN_TEMPLATE_CONTENT)
	if err != nil {
		return nil, 500, err
	}

	result = &AttemptSetIDEmailResult{
		UserID: thePath.UserID,
		Email:  theParams.Email,
	}

	return result, 200, nil
}
