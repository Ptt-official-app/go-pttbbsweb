package db

import (
	"time"

	"github.com/Ptt-official-app/go-openbbsmiddleware/config_util"
)

const configPrefix = "db"

func InitConfig() error {
	config()
	return nil
}

func setDurationConfig(idx string, orig time.Duration) time.Duration {
	return config_util.SetDurationConfig(configPrefix, idx, orig)
}
