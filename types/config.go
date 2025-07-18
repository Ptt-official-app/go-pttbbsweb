package types

import (
	"github.com/Ptt-official-app/go-pttbbs/bbs"
)

func config() {
	SERVICE_MODE = ServiceMode(setStringConfig("SERVICE_MODE", string(SERVICE_MODE)))

	HTTP_HOST = setStringConfig("HTTP_HOST", HTTP_HOST)
	URL_PREFIX = setStringConfig("URL_PREFIX", URL_PREFIX)
	BACKEND_PREFIX = setStringConfig("BACKEND_PREFIX", BACKEND_PREFIX)
	FRONTEND_PREFIX = setStringConfig("FRONTEND_PREFIX", FRONTEND_PREFIX)
	API_PREFIX = setStringConfig("API_PREFIX", API_PREFIX)

	PTTSYSOP = bbs.UUserID(setStringConfig("PTTSYSOP", string(PTTSYSOP)))

	BBSNAME = setStringConfig("BBSNAME", BBSNAME)
	BBSENAME = setStringConfig("BBSENAME", BBSENAME)

	// web
	STATIC_DIR = setStringConfig("STATIC_DIR", STATIC_DIR)

	ALLOW_ORIGINS = setListStringConfig("ALLOW_ORIGINS", ALLOW_ORIGINS)
	BLOCKED_REFERERS = setListStringConfig("BLOCKED_REFERERS", BLOCKED_REFERERS)
	IS_ALLOW_CROSSDOMAIN = setBoolConfig("IS_ALLOW_CROSSDOMAIN", IS_ALLOW_CROSSDOMAIN)

	COOKIE_DOMAIN = setStringConfig("COOKIE_DOMAIN", COOKIE_DOMAIN)
	TOKEN_COOKIE_SUFFIX = setStringConfig("TOKEN_COOKIE_SUFFIX", TOKEN_COOKIE_SUFFIX)

	CSRF_SECRET = setBytesConfig("CSRF_SECRET", CSRF_SECRET)
	CSRF_TOKEN = setStringConfig("CSRF_TOKEN", CSRF_TOKEN)
	CSRF_TOKEN_TS = setIntConfig("CSRF_TOKEN_TS", CSRF_TOKEN_TS)

	ACCESS_TOKEN_NAME = setStringConfig("ACCESS_TOKEN", ACCESS_TOKEN_NAME)
	ACCESS_TOKEN_EXPIRE_TS = setIntConfig("ACCESS_TOKEN_EXPIRE_TS", ACCESS_TOKEN_EXPIRE_TS)

	// email
	EMAIL_TOKEN_NAME = setStringConfig("EMAIL_TOKEN_NAME", EMAIL_TOKEN_NAME)

	EMAIL_FROM = setStringConfig("EMAIL_FROM", EMAIL_FROM)
	EMAIL_SERVER = setStringConfig("EMAIL_SERVER", EMAIL_SERVER)

	EMAILTOKEN_TEMPLATE = setStringConfig("EMAILTOKEN_TEMPLATE", EMAILTOKEN_TEMPLATE)
	IDEMAILTOKEN_TEMPLATE = setStringConfig("IDEMAILTOKEN_TEMPLATE", IDEMAILTOKEN_TEMPLATE)
	ATTEMPT_REGISTER_USER_TEMPLATE = setStringConfig("ATTEMPT_REGISTER_USER_TEMPLATE", ATTEMPT_REGISTER_USER_TEMPLATE)

	EXPIRE_USER_ID_EMAIL_IS_SET_NANO_TS = NanoTS(setInt64Config("EXPIRE_USER_ID_EMAIL_IS_SET_NANO_TS", int64(EXPIRE_USER_ID_EMAIL_IS_SET_NANO_TS)))
	EXPIRE_USER_EMAIL_IS_SET_NANO_TS = NanoTS(setInt64Config("EXPIRE_USER_EMAIL_IS_SET_NANO_TS", int64(EXPIRE_USER_EMAIL_IS_SET_NANO_TS)))

	EXPIRE_USER_ID_EMAIL_IS_NOT_SET_NANO_TS = NanoTS(setInt64Config("EXPIRE_USER_ID_EMAIL_IS_NOT_SET_NANO_TS", int64(EXPIRE_USER_ID_EMAIL_IS_NOT_SET_NANO_TS)))
	EXPIRE_USER_EMAIL_IS_NOT_SET_NANO_TS = NanoTS(setInt64Config("EXPIRE_USER_EMAIL_IS_NOT_SET_NANO_TS", int64(EXPIRE_USER_EMAIL_IS_NOT_SET_NANO_TS)))

	EXPIRE_ATTEMPT_REGISTER_USER_EMAIL_TS = setIntConfig("EXPIRE_ATTEMPT_REGISTER_USER_EMAIL_TS", EXPIRE_ATTEMPT_REGISTER_USER_EMAIL_TS)

	IS_2FA = setBoolConfig("IS_2FA", IS_2FA)
	MAX_2FA_TOKEN = setInt64Config("MAX_2FA_TOKEN", MAX_2FA_TOKEN)

	// big5
	BIG5_TO_UTF8 = setStringConfig("BIG5_TO_UTF8", BIG5_TO_UTF8)
	UTF8_TO_BIG5 = setStringConfig("UTF8_TO_BIG5", UTF8_TO_BIG5)
	AMBCJK = setStringConfig("AMBCJK", AMBCJK)

	// time-location
	TIME_LOCATION = setStringConfig("TIME_LOCATION", TIME_LOCATION)

	// carriage-return
	IS_CARRIAGE_RETURN = setBoolConfig("IS_CARRIAGE_RETURN", IS_CARRIAGE_RETURN)

	// is-all-guest
	IS_ALL_GUEST = setBoolConfig("IS_ALL_GUEST", IS_ALL_GUEST)

	// pttweb-hotboard-url
	PTTWEB_HOTBOARD_URL = setStringConfig("PTTWEB_HOTBOARD_URL", PTTWEB_HOTBOARD_URL)

	// expire-http-request-ts
	EXPIRE_HTTP_REQUEST_TS = setIntConfig("EXPIRE_HTTP_REQUEST_TS", EXPIRE_HTTP_REQUEST_TS)
}
