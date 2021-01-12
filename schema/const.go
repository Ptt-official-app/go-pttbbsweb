package schema

import (
	"github.com/Ptt-official-app/go-openbbsmiddleware/db"
	redis "github.com/go-redis/redis/v8"
)

var (
	client *db.Client

	rdb *redis.Client
)

const ()
