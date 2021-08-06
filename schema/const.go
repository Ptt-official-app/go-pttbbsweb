package schema

import (
	"github.com/Ptt-official-app/go-openbbsmiddleware/db"
	redis "github.com/go-redis/redis/v8"
)

const (
	TITLE_REGEX_N_GRAM = 5
)

var (
	client *db.Client

	rdb *redis.Client
)
