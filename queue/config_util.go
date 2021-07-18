package queue

import (
	"github.com/Ptt-official-app/go-openbbsmiddleware/configutil"
)

const configPrefix = "go-openbbsmiddleware:queue"

func InitConfig() error {
	config()
	return nil
}

func setIntConfig(idx string, orig int) int {
	return configutil.SetIntConfig(configPrefix, idx, orig)
}
