package types

import (
	"time"

	"github.com/Ptt-official-app/go-pttbbs/bbs"
)

var (
	SERVICE_MODE = DEV //can be DEV, PRODUCTION, INFO, DEBUG

	HTTP_HOST      = "localhost:3457"            //serving http-host
	URL_PREFIX     = "http://localhost:3457/bbs" //advertising url-prefix
	BACKEND_PREFIX = "http://localhost:3456/v1"  //backend url-prefix

	PTTSYSOP = bbs.UUserID("SYSOP")

	//web
	STATIC_DIR           = "static"
	ALLOW_ORIGINS        = []string{}
	ALLOW_ORIGINS_MAP    = map[string]bool{}
	BLOCKED_REFERERS     = []string{}
	BLOCKED_REFERERS_MAP = map[string]bool{}
	IS_ALLOW_CROSSDOMAIN = true

	COOKIE_DOMAIN       = "localhost"
	TOKEN_COOKIE_SUFFIX = "Secure;"

	CSRF_SECRET            = []byte("test_csrf_secret")
	CSRF_TOKEN             = "csrftoken"
	CSRF_TOKEN_TS          = 0
	CSRF_TOKEN_TS_DURATION = time.Duration(CSRF_TOKEN_TS) * time.Second

	ACCESS_TOKEN                    = "token"
	ACCESS_TOKEN_EXPIRE_TS          = 86400
	ACCESS_TOKEN_EXPIRE_TS_DURATION = time.Duration(ACCESS_TOKEN_EXPIRE_TS) * time.Second

	//big5
	BIG5_TO_UTF8 = "types/uao250-b2u.big5.txt"
	UTF8_TO_BIG5 = "types/uao250-u2b.big5.txt"
	AMBCJK       = "types/ambcjk.big5.txt"

	//time-location
	TIME_LOCATION = "Asia/Taipei"
)
