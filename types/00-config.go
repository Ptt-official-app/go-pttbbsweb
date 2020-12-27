package types

var (
	SERVICE_MODE = DEV //can be DEV, PRODUCTION, INFO, DEBUG

	HTTP_HOST      = "localhost:3457"            //serving http-host
	URL_PREFIX     = "http://localhost:3457/bbs" //advertising url-prefix
	BACKEND_PREFIX = "http://localhost:3456/v1"  //backend url-prefix

	//big5
	BIG5_TO_UTF8 = "types/uao250-b2u.big5.txt"
	UTF8_TO_BIG5 = "types/uao250-u2b.big5.txt"
	AMBCJK       = "types/ambcjk.big5.txt"

	TIME_LOCATION = "Asia/Taipei"
)
