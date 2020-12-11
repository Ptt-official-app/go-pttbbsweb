package api

import "github.com/Ptt-official-app/go-openbbsmiddleware/config_util"

const configPrefix = "api"

func InitConfig() error {
	config()
	return nil
}

func setStringConfig(idx string, orig string) string {
	return config_util.SetStringConfig(configPrefix, idx, orig)
}

func setBytesConfig(idx string, orig []byte) []byte {
	return config_util.SetBytesConfig(configPrefix, idx, orig)
}
