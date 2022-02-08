package dbcs

import (
	"github.com/Ptt-official-app/go-openbbsmiddleware/configutil"
	"github.com/Ptt-official-app/go-openbbsmiddleware/types"
)

const configPrefix = "go-openbbsmiddleware:dbcs"

func InitConfig() (err error) {
	config()

	return nil
}

func setNanoTSConfig(idx string, orig types.NanoTS) types.NanoTS {
	return types.NanoTS(configutil.SetInt64Config(configPrefix, idx, int64(orig)))
}
