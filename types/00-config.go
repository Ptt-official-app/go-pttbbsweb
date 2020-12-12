package types

var (
	SERVICE_MODE = DEV //can be DEV, PRODUCTION, INFO, DEBUG

	HTTP_HOST = "localhost:3457"

	//Creating JWT Token
	JWT_SECRET = []byte("jwt_secret")
	JWT_ISSUER = "go-pttbbs"
)
