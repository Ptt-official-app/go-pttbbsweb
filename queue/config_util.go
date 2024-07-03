package queue

import (
	"github.com/Ptt-official-app/go-pttbbsweb/configutil"
)

const configPrefix = "go-pttbbsweb:queue"

func InitConfig() error {
	config()
	return nil
}

func setIntConfig(idx string, orig int) int {
	return configutil.SetIntConfig(configPrefix, idx, orig)
}
