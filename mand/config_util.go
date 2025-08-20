package mand

import "github.com/Ptt-official-app/pttbbs-backend/configutil"

const configPrefix = "pttbbs-backend:mand"

func InitConfig() error {
	config()
	return nil
}

func setStringConfig(idx string, orig string) string {
	return configutil.SetStringConfig(configPrefix, idx, orig)
}
