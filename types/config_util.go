package types

import (
	"time"

	"github.com/Ptt-official-app/go-openbbsmiddleware/config_util"
)

const configPrefix = "go-openbbsmiddleware:types"

func InitConfig() (err error) {
	config()

	postConfig()

	err = initBig5()
	if err != nil {
		return err
	}
	return nil
}

func setStringConfig(idx string, orig string) string {
	return config_util.SetStringConfig(configPrefix, idx, orig)
}

func setListStringConfig(idx string, orig []string) []string {
	return config_util.SetListStringConfig(configPrefix, idx, orig)
}

func setBytesConfig(idx string, orig []byte) []byte {
	return config_util.SetBytesConfig(configPrefix, idx, orig)
}

func setBoolConfig(idx string, orig bool) bool {
	return config_util.SetBoolConfig(configPrefix, idx, orig)
}

func setIntConfig(idx string, orig int) int {
	return config_util.SetIntConfig(configPrefix, idx, orig)
}

func postConfig() {
	setTimeLocation(TIME_LOCATION)
	setAllowOrigins(ALLOW_ORIGINS)
	setBlockedReferers(BLOCKED_REFERERS)
	setCSRFTokenTS(CSRF_TOKEN_TS)
	setAccessTokenExpireTS(ACCESS_TOKEN_EXPIRE_TS)
}

//setTimeLocation
//
//
func setTimeLocation(timeLocation string) (origTimeLocation string, err error) {
	origTimeLocation = TIME_LOCATION
	TIME_LOCATION = timeLocation

	TIMEZONE, err = time.LoadLocation(TIME_LOCATION)

	return origTimeLocation, err
}

func setAllowOrigins(allowOrigins []string) (origAllowOrigins []string, err error) {

	origAllowOrigins = ALLOW_ORIGINS

	ALLOW_ORIGINS = allowOrigins
	newAllowOriginsMap := map[string]bool{}

	for _, each := range allowOrigins {
		newAllowOriginsMap[each] = true
	}

	ALLOW_ORIGINS_MAP = newAllowOriginsMap

	return origAllowOrigins, nil
}

func setBlockedReferers(blockedReferers []string) (origBlockedReferers []string, err error) {
	origBlockedReferers = BLOCKED_REFERERS

	BLOCKED_REFERERS = blockedReferers
	newBlockedReferersMap := map[string]bool{}

	for _, each := range blockedReferers {
		newBlockedReferersMap[each] = true
	}

	BLOCKED_REFERERS_MAP = newBlockedReferersMap

	return origBlockedReferers, nil
}

func setCSRFTokenTS(csrfTokenTS int) (origCSRFTokenTS int, err error) {
	origCSRFTokenTS = CSRF_TOKEN_TS

	CSRF_TOKEN_TS = csrfTokenTS

	CSRF_TOKEN_TS_DURATION = time.Duration(CSRF_TOKEN_TS) * time.Second

	return origCSRFTokenTS, nil
}

func setAccessTokenExpireTS(accessTokenExpireTS int) (origAccessTokenExpireTS int, err error) {

	origAccessTokenExpireTS = ACCESS_TOKEN_EXPIRE_TS

	ACCESS_TOKEN_EXPIRE_TS = accessTokenExpireTS

	ACCESS_TOKEN_EXPIRE_TS_DURATION = time.Duration(ACCESS_TOKEN_EXPIRE_TS) * time.Second

	return origAccessTokenExpireTS, nil

}
