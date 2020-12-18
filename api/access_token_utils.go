package api

import (
	"github.com/Ptt-official-app/go-openbbsmiddleware/schema"
	pttbbsapi "github.com/Ptt-official-app/go-pttbbs/api"
	"github.com/Ptt-official-app/go-pttbbs/bbs"
)

func serializeAccessTokenAndUpdateDB(userID bbs.UUserID, jwt string) (accessToken *schema.AccessToken, err error) {

	accessToken = schema.NewAccessToken(userID, jwt)

	if string(userID) == pttbbsapi.GUEST {
		return accessToken, nil
	}

	err = schema.UpdateAccessToken(accessToken)
	if err != nil {
		return nil, err
	}

	return accessToken, nil
}
