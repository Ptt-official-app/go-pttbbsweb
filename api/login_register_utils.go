package api

import (
	"fmt"
	"strings"

	"github.com/Ptt-official-app/go-openbbsmiddleware/schema"
	"github.com/Ptt-official-app/go-openbbsmiddleware/types"
	"github.com/Ptt-official-app/go-openbbsmiddleware/utils"
	"github.com/Ptt-official-app/go-pttbbs/bbs"
	"github.com/gin-gonic/gin"
)

func setTokenToCookie(c *gin.Context, accessToken string) {
	setCookie(c, types.ACCESS_TOKEN_NAME, accessToken, types.ACCESS_TOKEN_EXPIRE_TS_DURATION, true)
}

func gen2FATokenAndSendEmail(userID bbs.UUserID, email string, title string, template string) (err error) {
	token := gen2FAToken()

	err = schema.Set2FA(userID, email, token, types.EXPIRE_ATTEMPT_REGISTER_USER_EMAIL_TS_DURATION)
	if err != nil {
		return err
	}

	content := strings.Replace(
		strings.Replace(
			template, "__USER__", string(userID), -1,
		), "__TOKEN__", token, -1,
	)

	return utils.SendEmail([]string{email}, title, content)
}

func gen2FAToken() string {
	randInt := utils.GenRandomInt64(types.MAX_2FA_TOKEN)
	return fmt.Sprintf("%06d", randInt)
}

func check2FAToken(userID bbs.UUserID, email string, token string) (err error) {
	emailtoken_db, err := schema.Get2FA(userID)
	if err != nil {
		return err
	}

	emailtoken := schema.TwoFactorSerializeValue(email, token)

	if emailtoken != emailtoken_db {
		return ErrInvalidToken
	}

	return nil
}
