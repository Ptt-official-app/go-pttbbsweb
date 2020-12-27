package mock_http

import (
	"time"

	"github.com/Ptt-official-app/go-pttbbs/api"
	"github.com/Ptt-official-app/go-pttbbs/bbs"
	"gopkg.in/square/go-jose.v2"
	"gopkg.in/square/go-jose.v2/jwt"
)

func Login(params *api.LoginParams) (ret *api.LoginResult) {
	userID := params.Username
	token, _ := createToken(userID)

	ret = &api.LoginResult{
		UserID:    bbs.UUserID(userID),
		Jwt:       token,
		TokenType: "bearer",
	}

	return ret
}

func createToken(userID string) (string, error) {
	var err error

	sig, err := jose.NewSigner(jose.SigningKey{Algorithm: jose.HS256, Key: api.JWT_SECRET}, (&jose.SignerOptions{}).WithType("JWT"))
	if err != nil {
		return "", err
	}

	cl := &api.JwtClaim{
		UUserID: bbs.UUserID(userID),
		Expire:  jwt.NewNumericDate(time.Now().Add(time.Hour * 72)),
	}

	raw, err := jwt.Signed(sig).Claims(cl).CompactSerialize()
	if err != nil {
		return "", err
	}

	return raw, nil
}
