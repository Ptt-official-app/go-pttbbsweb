package types

func config() {
	SERVICE_MODE = ServiceMode(setStringConfig("SERVICE_MODE", string(SERVICE_MODE)))

	HTTP_HOST = setStringConfig("HTTP_HOST", HTTP_HOST)
	URL_PREFIX = setStringConfig("URL_PREFIX", URL_PREFIX)
	BACKEND_PREFIX = setStringConfig("BACKEND_PREFIX", BACKEND_PREFIX)

	BIG5_TO_UTF8 = setStringConfig("BIG5_TO_UTF8", BIG5_TO_UTF8)
	UTF8_TO_BIG5 = setStringConfig("UTF8_TO_BIG5", UTF8_TO_BIG5)
	TIME_LOCATION = setStringConfig("TIME_LOCATION", TIME_LOCATION)
}
