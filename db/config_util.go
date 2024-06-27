package db

import (
	"time"

	"github.com/Ptt-official-app/go-pttbbsweb/configutil"
)

const configPrefix = "go-pttbbsweb:db"

func InitConfig() error {
	config()
	return nil
}

func setDurationConfig(idx string, orig time.Duration) time.Duration {
	return configutil.SetDurationConfig(configPrefix, idx, orig)
}
