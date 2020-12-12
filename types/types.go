package types

import "gopkg.in/square/go-jose.v2/jwt"

type JwtClaim struct {
	UserID string
	Expire *jwt.NumericDate
}
