package dbcs

import (
	"github.com/Ptt-official-app/pttbbs-backend/configutil"
	"github.com/Ptt-official-app/pttbbs-backend/types"
)

const configPrefix = "pttbbs-backend:dbcs"

func InitConfig() (err error) {
	config()

	return nil
}

func setNanoTSConfig(idx string, orig types.NanoTS) types.NanoTS {
	return types.NanoTS(configutil.SetInt64Config(configPrefix, idx, int64(orig)))
}
