package api

import (
	"fmt"
	"strings"

	"github.com/Ptt-official-app/go-pttbbs/bbs"
	"github.com/Ptt-official-app/pttbbs-backend/schema"
	"github.com/Ptt-official-app/pttbbs-backend/types"
	"github.com/Ptt-official-app/pttbbs-backend/utils"
	"github.com/gin-gonic/gin"
)

func setTokenToCookie(c *gin.Context, accessToken string) {
	setCookie(c, types.ACCESS_TOKEN_NAME, accessToken, types.ACCESS_TOKEN_EXPIRE_TS_DURATION, true)
}

func removeTokenFromCookie(c *gin.Context) {
	removeCookie(c, types.ACCESS_TOKEN_NAME, true)
}

func gen2FATokenAndSendEmail(userID bbs.UUserID, email string, title string, template string) (err error) {
	token := gen2FAToken()

	err = schema.Set2FA(userID, email, token, types.EXPIRE_ATTEMPT_REGISTER_USER_EMAIL_TS_DURATION)
	if err != nil {
		return err
	}

	content := strings.ReplaceAll(
		strings.ReplaceAll(
			template, "__USER__", string(userID),
		), "__TOKEN__", token,
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
