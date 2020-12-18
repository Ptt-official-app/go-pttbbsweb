package main

import (
	"time"

	"github.com/Ptt-official-app/go-openbbsmiddleware/config_util"
)

const configPrefix = "go-openbbsmiddleware"

func InitConfig() error {
	config()
	return nil
}

func setStringConfig(idx string, orig string) string {
	return config_util.SetStringConfig(configPrefix, idx, orig)
}

func setDurationConfig(idx string, orig time.Duration) time.Duration {
	return config_util.SetDurationConfig(configPrefix, idx, orig)
}
