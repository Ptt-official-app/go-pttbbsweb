package mock_http

import "github.com/Ptt-official-app/go-pttbbs/api"

func GetEmailTokenInfo(params *api.GetEmailTokenInfoParams) (ret *api.GetEmailTokenInfoResult) {

	userID, clientInfo, email, _ := api.VerifyEmailJwt(params.Jwt, params.Context)

	return &api.GetEmailTokenInfoResult{UserID: userID, ClientInfo: clientInfo, Email: email}
}
