package api

import (
	"fmt"
	"strings"

	pttbbsapi "github.com/Ptt-official-app/go-pttbbs/api"
	"github.com/Ptt-official-app/go-pttbbs/bbs"
	"github.com/Ptt-official-app/go-pttbbsweb/schema"
	"github.com/Ptt-official-app/go-pttbbsweb/types"
	"github.com/Ptt-official-app/go-pttbbsweb/utils"
	"github.com/gin-gonic/gin"
)

func deserializeUserDetailAndUpdateDBGuest(updateNanoTS types.NanoTS) (userDetail *schema.UserDetail, err error) {
	userDetail = schema.NewUserDetailGuest(updateNanoTS)

	err = schema.UpdateUserDetail(userDetail)
	if err != nil {
		return nil, err
	}

	return userDetail, nil
}

func deserializeUserDetailAndUpdateDB(user_b pttbbsapi.GetUserResult, updateNanoTS types.NanoTS) (userDetail *schema.UserDetail, err error) {
	userDetail = schema.NewUserDetail(user_b, updateNanoTS)

	err = schema.UpdateUserDetail(userDetail)
	if err != nil {
		return nil, err
	}

	return userDetail, nil
}

func deserializeEmailTokenAndEmail(email string, title string, userID bbs.UUserID, jwt string, urlTemplate string, contentTemplate string) (err error) {
	content := deserializeEmailToken(email, userID, jwt, urlTemplate, contentTemplate)

	return utils.SendEmail([]string{email}, title, content)
}

func deserializeEmailToken(email string, userID bbs.UUserID, token string, urlTemplate string, contentTemplate string) (content string) {
	userIDStr := string(userID)

	urlMap := map[string]string{
		"user_id": userIDStr,
	}
	url := types.FRONTEND_PREFIX + utils.MergeURL(urlMap, urlTemplate)

	url += fmt.Sprintf("?%v=%v", types.EMAIL_TOKEN_NAME, token)

	content = strings.ReplaceAll(
		strings.ReplaceAll(
			strings.ReplaceAll(
				contentTemplate, "__USER__", userIDStr,
			), "__URL__", url,
		), "__EMAIL__", email,
	)

	return content
}

func checkUniqueIDEmail(email string) (err error) {
	nowNanoTS := types.NowNanoTS()
	userIDEmail, err := schema.GetUserIDEmailByEmail(email, nowNanoTS)
	if err != nil {
		return err
	}

	if userIDEmail != nil {
		return ErrAlreadyExists
	}

	return nil
}

func checkUniqueEmail(email string) (err error) {
	nowNanoTS := types.NowNanoTS()
	userEmail, err := schema.GetUserEmailByEmail(email, nowNanoTS)
	if err != nil {
		return err
	}

	if userEmail != nil {
		return ErrAlreadyExists
	}

	return nil
}

func getEmailTokenInfo(jwt string, context pttbbsapi.EmailTokenContext, c *gin.Context) (userID bbs.UUserID, email string, clientInfo string, statusCode int, err error) {
	theParams_b := &pttbbsapi.GetEmailTokenInfoParams{
		Jwt:     jwt,
		Context: context,
	}

	var result_b *pttbbsapi.GetEmailTokenInfoResult

	statusCode, err = utils.BackendPost(c, pttbbsapi.GET_EMAIL_TOKEN_INFO_R, theParams_b, nil, &result_b)
	if err != nil {
		return "", "", "", statusCode, err
	}

	return result_b.UserID, result_b.Email, result_b.ClientInfo, statusCode, err
}
