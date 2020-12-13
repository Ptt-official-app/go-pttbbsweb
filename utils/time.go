package utils

import (
	"time"

	"github.com/Ptt-official-app/go-openbbsmiddleware/types"
)

func GetNowNanoTS() types.NanoTS {
	return types.NanoTS(time.Now().UnixNano())
}
