package api

func config() {
	JWT_SECRET = setBytesConfig("JWT_SECRET", JWT_SECRET)
	JWT_ISSUER = setStringConfig("JWT_ISSUER", JWT_ISSUER)
}
