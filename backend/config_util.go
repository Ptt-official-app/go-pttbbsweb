package backend

import "github.com/Ptt-official-app/go-openbbsmiddleware/config_util"

const configPrefix = "backend"

func InitConfig() error {
	config()
	return nil
}

func setStringConfig(idx string, orig string) string {
	return config_util.SetStringConfig(configPrefix, idx, orig)
}
