package db

import (
	"time"

	"github.com/Ptt-official-app/pttbbs-backend/configutil"
)

const configPrefix = "pttbbs-backend:db"

func InitConfig() error {
	config()
	return nil
}

func setDurationConfig(idx string, orig time.Duration) time.Duration {
	return configutil.SetDurationConfig(configPrefix, idx, orig)
}
