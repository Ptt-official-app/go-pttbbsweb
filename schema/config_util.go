package schema

import (
	"github.com/Ptt-official-app/go-openbbsmiddleware/config_util"
)

const configPrefix = "schema"

func InitConfig() error {
	config()
	return nil
}

func setStringConfig(idx string, orig string) string {
	return config_util.SetStringConfig(configPrefix, idx, orig)
}

func setIntConfig(idx string, orig int) int {
	return config_util.SetIntConfig(configPrefix, idx, orig)
}
