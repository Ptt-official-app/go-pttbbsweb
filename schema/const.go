package schema

import (
	"time"

	"github.com/Ptt-official-app/go-openbbsmiddleware/db"
	redis "github.com/go-redis/redis/v8"
)

const (
	TITLE_REGEX_N_GRAM              = 5
	TIME_CALC_ALL_USER_VISIT_COUNTS = -10 * time.Minute

	MAX_CONTENT_BLOCK = 5

	MAX_ALL_CONTENT_BLOCK = 2000
)

var (
	client *db.Client

	rdb *redis.Client
)
