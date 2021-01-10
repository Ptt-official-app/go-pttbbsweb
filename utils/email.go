package utils

import (
	"net/smtp"

	"github.com/Ptt-official-app/go-openbbsmiddleware/types"
)

func SendEmail(rcpts []string, content string) (err error) {
	if isTest {
		return nil
	}

	from := types.EMAIL_FROM
	username := ""
	password := ""

	auth := smtp.PlainAuth("", username, password, types.EMAIL_SERVER)

	return smtp.SendMail(types.EMAIL_SERVER, auth, from, rcpts, []byte(content))
}
