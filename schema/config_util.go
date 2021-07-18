package schema

import (
	"time"

	"github.com/Ptt-official-app/go-openbbsmiddleware/configutil"
)

const configPrefix = "go-openbbsmiddleware:schema"

func InitConfig() error {
	config()
	return nil
}

func setStringConfig(idx string, orig string) string {
	return configutil.SetStringConfig(configPrefix, idx, orig)
}

func setIntConfig(idx string, orig int) int {
	return configutil.SetIntConfig(configPrefix, idx, orig)
}

func setDurationConfig(idx string, orig time.Duration) time.Duration {
	return configutil.SetDurationConfig(configPrefix, idx, orig)
}
