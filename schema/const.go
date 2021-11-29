package schema

import (
	"time"

	"github.com/Ptt-official-app/go-openbbsmiddleware/db"
	redis "github.com/go-redis/redis/v8"
)

const (
	TITLE_REGEX_N_GRAM              = 5
	TIME_CALC_ALL_USER_VISIT_COUNTS = -10 * time.Minute
)

var (
	client *db.Client

	rdb *redis.Client
)
