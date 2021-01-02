package types

import "github.com/Ptt-official-app/go-pttbbs/bbs"

func config() {
	SERVICE_MODE = ServiceMode(setStringConfig("SERVICE_MODE", string(SERVICE_MODE)))

	HTTP_HOST = setStringConfig("HTTP_HOST", HTTP_HOST)
	URL_PREFIX = setStringConfig("URL_PREFIX", URL_PREFIX)
	BACKEND_PREFIX = setStringConfig("BACKEND_PREFIX", BACKEND_PREFIX)

	PTTSYSOP = bbs.UUserID(setStringConfig("PTTSYSOP", string(PTTSYSOP)))

	STATIC_DIR = setStringConfig("STATIC_DIR", STATIC_DIR)
	CSRF_SECRET = setBytesConfig("CSRF_SECRET", CSRF_SECRET)
	ALLOW_ORIGINS = setListStringConfig("ALLOW_ORIGINS", ALLOW_ORIGINS)
	BLOCKED_REFERERS = setListStringConfig("BLOCKED_REFERERS", BLOCKED_REFERERS)
	IS_ALLOW_CROSSDOMAIN = setBoolConfig("IS_ALLOW_CROSSDOMAIN", IS_ALLOW_CROSSDOMAIN)

	BIG5_TO_UTF8 = setStringConfig("BIG5_TO_UTF8", BIG5_TO_UTF8)
	UTF8_TO_BIG5 = setStringConfig("UTF8_TO_BIG5", UTF8_TO_BIG5)
	TIME_LOCATION = setStringConfig("TIME_LOCATION", TIME_LOCATION)
}
