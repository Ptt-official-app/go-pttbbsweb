package api

import (
	"time"

	"github.com/Ptt-official-app/go-openbbsmiddleware/schema"
	"github.com/Ptt-official-app/go-openbbsmiddleware/types"
	pttbbsapi "github.com/Ptt-official-app/go-pttbbs/api"
	"github.com/Ptt-official-app/go-pttbbs/bbs"
	"gopkg.in/square/go-jose.v2/jwt"
)

func isValidClient(clientID string, clientSecret string) bool {

	query := &schema.RegisterClientQuery{
		ClientID: clientID,
	}

	ret := &schema.Client{}
	err := schema.Client_c.FindOne(query, ret, nil)
	if err != nil {
		return false
	}

	return ret.ClientSecret == clientSecret
}

//verifyJwt
//
//from https://github.com/Ptt-official-app/go-pttbbs/blob/main/api.go#L93
func VerifyJwt(raw string) (userID bbs.UUserID, err error) {
	tok, err := jwt.ParseSigned(raw)
	if err != nil {
		return "", ErrInvalidToken
	}

	cl := &pttbbsapi.JwtClaim{}
	if err := tok.Claims(types.JWT_SECRET, cl); err != nil {
		return "", ErrInvalidToken
	}

	currentNanoTS := jwt.NewNumericDate(time.Now())
	if *currentNanoTS > *cl.Expire {
		return "", ErrInvalidToken
	}

	return cl.UUserID, nil
}
