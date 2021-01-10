package schema

import (
	"github.com/Ptt-official-app/go-openbbsmiddleware/db"
	redis "github.com/go-redis/redis/v8"
)

var (
	client *db.Client

	rdb *redis.Client
)

const (
	EXPIRE_USER_ID_EMAIL_SET_NANO_TS = 10 * 1000000000 //5 seconds
)
