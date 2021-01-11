package types

import (
	"time"

	"github.com/Ptt-official-app/go-pttbbs/bbs"
)

var (
	SERVICE_MODE = DEV //can be DEV, PRODUCTION, INFO, DEBUG

	HTTP_HOST       = "localhost:3457"            //serving http-host
	URL_PREFIX      = "http://localhost:3457/bbs" //advertising url-prefix
	BACKEND_PREFIX  = "http://localhost:3456/v1"  //backend url-prefix
	FRONTEND_PREFIX = "http://localhost:3457"     //frontend-prefix, email
	API_PREFIX      = "/api"                      //api-prefix

	PTTSYSOP = bbs.UUserID("SYSOP")

	BBSNAME  = "新批踢踢" /* 中文站名 */
	BBSENAME = "PTT2" /* 英文站名 */

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

	ACCESS_TOKEN_NAME               = "token" //access-token-name in cookie
	ACCESS_TOKEN_EXPIRE_TS          = 86400
	ACCESS_TOKEN_EXPIRE_TS_DURATION = time.Duration(ACCESS_TOKEN_EXPIRE_TS) * time.Second

	//email
	EMAIL_TOKEN_NAME = "token" //email-token in email-url

	EMAIL_FROM   = "noreply@localhost"
	EMAIL_SERVER = "localhost:25"

	EMAILTOKEN_TITLE      = "更換 " + BBSNAME + " 的聯絡信箱 (Update " + BBSENAME + " Contact Email)"
	IDEMAILTOKEN_TITLE    = "更換 " + BBSNAME + " 的認證信箱 (Update " + BBSENAME + " Identity Email)"
	EMAILTOKEN_TEMPLATE   = "/etc/go-openbbsmiddleware/emailtoken.template"
	IDEMAILTOKEN_TEMPLATE = "/etc/go-openbbsmiddleware/idemailtoken.template"

	EMAILTOKEN_TEMPLATE_CONTENT   = "__EMAIL__, __USER__, __URL__"
	IDEMAILTOKEN_TEMPLATE_CONTENT = "__EMAIL__, __USER__, __URL__"

	EXPIRE_USER_ID_EMAIL_IS_SET_NANO_TS = NanoTS(100 * 86400 * 1000000000) //100 days
	EXPIRE_USER_EMAIL_IS_SET_NANO_TS    = NanoTS(1 * 86400 * 1000000000)   //1 day

	EXPIRE_USER_ID_EMAIL_IS_NOT_SET_NANO_TS = NanoTS(900 * 1000000000) //15 mins
	EXPIRE_USER_EMAIL_IS_NOT_SET_NANO_TS    = NanoTS(900 * 1000000000) //15 mins

	//big5
	BIG5_TO_UTF8 = "types/uao250-b2u.big5.txt"
	UTF8_TO_BIG5 = "types/uao250-u2b.big5.txt"
	AMBCJK       = "types/ambcjk.big5.txt"

	//time-location
	TIME_LOCATION = "Asia/Taipei"
)
