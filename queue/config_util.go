package queue

import (
	"github.com/Ptt-official-app/pttbbs-backend/configutil"
)

const configPrefix = "pttbbs-backend:queue"

func InitConfig() error {
	config()
	return nil
}

func setIntConfig(idx string, orig int) int {
	return configutil.SetIntConfig(configPrefix, idx, orig)
}
