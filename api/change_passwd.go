package api

import (
	"github.com/Ptt-official-app/go-openbbsmiddleware/schema"
	"github.com/Ptt-official-app/go-openbbsmiddleware/types"
	"github.com/Ptt-official-app/go-openbbsmiddleware/utils"
	pttbbsapi "github.com/Ptt-official-app/go-pttbbs/api"
	"github.com/Ptt-official-app/go-pttbbs/bbs"
	"github.com/gin-gonic/gin"
)

const CHANGE_PASSWD_R = "/user/:user_id/updatepasswd"

type ChangePasswdParams struct {
	ClientID     string `json:"client_id" form:"client_id"`
	ClientSecret string `json:"client_secret" form:"client_secret"`

	OrigPassword    string `json:"orig_password" form:"orig_password"`
	Password        string `json:"password" form:"password"`
	PasswordConfirm string `json:"password_confirm" form:"password_confirm"`
}

type ChangePasswdPath struct {
	UserID bbs.UUserID `uri:"user_id"`
}

type ChangePasswdResult struct {
	UserID      bbs.UUserID `json:"user_id"`
	AccessToken string      `json:"access_token"`
	TokenType   string      `json:"token_type"`
}

func ChangePasswdWrapper(c *gin.Context) {
	params := &ChangePasswdParams{}
	path := &ChangePasswdPath{}
	LoginRequiredPathJSON(ChangePasswd, params, path, c)
}

func ChangePasswd(remoteAddr string, userID bbs.UUserID, params interface{}, path interface{}, c *gin.Context) (result interface{}, statusCode int, err error) {
	theParams, ok := params.(*ChangePasswdParams)
	if !ok {
		return nil, 400, ErrInvalidParams
	}

	thePath, ok := path.(*ChangePasswdPath)
	if !ok {
		return nil, 400, ErrInvalidPath
	}

	if theParams.Password != theParams.PasswordConfirm {
		return nil, 400, ErrInvalidParams
	}

	isValidClient, client := checkClient(theParams.ClientID, theParams.ClientSecret)

	if !isValidClient {
		return nil, 400, ErrInvalidParams
	}

	clientInfo := getClientInfo(client)

	// get backend data
	theParams_b := &pttbbsapi.ChangePasswdParams{
		ClientInfo: clientInfo,
		OrigPasswd: theParams.OrigPassword,
		Passwd:     theParams.Password,
	}

	var result_b *pttbbsapi.ChangePasswdResult

	urlMap := make(map[string]string)
	urlMap["uid"] = string(thePath.UserID)
	url := utils.MergeURL(urlMap, pttbbsapi.CHANGE_PASSWD_R)

	statusCode, err = utils.BackendPost(c, url, theParams_b, nil, &result_b)
	if err != nil {
		return nil, statusCode, err
	}

	// update db
	updateNanoTS := types.NowNanoTS()
	accessToken_db, err := deserializeAccessTokenAndUpdateDB(result_b.UserID, result_b.Jwt, updateNanoTS)
	if err != nil {
		return nil, 500, err
	}
	// result
	result = NewChangePasswdResult(accessToken_db)

	setTokenToCookie(c, accessToken_db.AccessToken)

	return result, 200, nil
}

func NewChangePasswdResult(accessToken_db *schema.AccessToken) *ChangePasswdResult {
	return &ChangePasswdResult{
		UserID:      accessToken_db.UserID,
		AccessToken: accessToken_db.AccessToken,
		TokenType:   "bearer",
	}
}
