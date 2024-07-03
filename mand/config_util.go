package mand

import "github.com/Ptt-official-app/go-pttbbsweb/configutil"

const configPrefix = "go-pttbbsweb:mand"

func InitConfig() error {
	config()
	return nil
}

func setStringConfig(idx string, orig string) string {
	return configutil.SetStringConfig(configPrefix, idx, orig)
}
