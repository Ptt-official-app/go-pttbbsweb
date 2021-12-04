package schema

import (
	"context"
	"time"
)

var (
	REDIS_USER_VISIT_COUNT     = "user_visit_count"
	REDIS_USER_VISIT_COUNT_EXP = 10 * time.Minute
)

// SetUserVisitCount try to set user visit count to redis
func SetUserVisitCount(count int64) error {
	ctx := context.Background()
	err := rdb.Set(ctx, REDIS_USER_VISIT_COUNT, count, REDIS_USER_VISIT_COUNT_EXP).Err()
	if err != nil {
		return err
	}
	return nil
}

func GetUserVisitCount() (total int64) {
	ctx := context.Background()
	val, err := rdb.Get(ctx, REDIS_USER_VISIT_COUNT).Int64()
	if err != nil {
		return 0
	}
	return val
}
