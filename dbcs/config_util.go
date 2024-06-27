package dbcs

import (
	"github.com/Ptt-official-app/go-pttbbsweb/configutil"
	"github.com/Ptt-official-app/go-pttbbsweb/types"
)

const configPrefix = "go-pttbbsweb:dbcs"

func InitConfig() (err error) {
	config()

	return nil
}

func setNanoTSConfig(idx string, orig types.NanoTS) types.NanoTS {
	return types.NanoTS(configutil.SetInt64Config(configPrefix, idx, int64(orig)))
}
