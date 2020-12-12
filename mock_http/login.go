package mock_http

import (
	"time"

	"github.com/Ptt-official-app/go-openbbsmiddleware/backend"
	"github.com/Ptt-official-app/go-openbbsmiddleware/types"
	"gopkg.in/square/go-jose.v2"
	"gopkg.in/square/go-jose.v2/jwt"
)

func Login(params *backend.LoginParams) (ret *backend.LoginResult) {
	userID := params.UserID
	token, _ := createToken(userID)

	ret = &backend.LoginResult{
		Jwt:       token,
		TokenType: "bearer",
	}

	return ret
}

func createToken(userID string) (string, error) {
	var err error

	sig, err := jose.NewSigner(jose.SigningKey{Algorithm: jose.HS256, Key: types.JWT_SECRET}, (&jose.SignerOptions{}).WithType("JWT"))
	if err != nil {
		return "", err
	}

	cl := &types.JwtClaim{
		UserID: userID,
		Expire: jwt.NewNumericDate(time.Now().Add(time.Hour * 72)),
	}

	raw, err := jwt.Signed(sig).Claims(cl).CompactSerialize()
	if err != nil {
		return "", err
	}

	return raw, nil
}
