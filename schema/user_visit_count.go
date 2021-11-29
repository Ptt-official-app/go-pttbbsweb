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
