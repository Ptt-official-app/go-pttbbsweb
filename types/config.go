package types

func config() {
	HTTP_HOST = setStringConfig("HTTP_HOST", HTTP_HOST)

	JWT_SECRET = setBytesConfig("JWT_SECRET", JWT_SECRET)
	JWT_ISSUER = setStringConfig("JWT_ISSUER", JWT_ISSUER)
}
