package schema

import (
	"context"
	"time"

	"github.com/Ptt-official-app/pttbbs-backend/types"
)

// TryLock
func TryLock(key string, expireTSDuration time.Duration) (err error) {
	ctx, cancel := context.WithTimeout(context.Background(), REDIS_TIMEOUT_MILLI_TS*time.Millisecond)
	defer func() {
		ctxErr := ctx.Err()
		cancel()
		if err == nil {
			err = ctxErr
		}
	}()

	updateNanoTS := int64(types.NowNanoTS())

	val, err := rdb.SetNX(ctx, "lock:"+key, updateNanoTS, expireTSDuration).Result()
	if err != nil {
		return err
	}
	if !val {
		return ErrNoLock
	}

	return nil
}

// Unlock
func Unlock(key string) (err error) {
	ctx, cancel := context.WithTimeout(context.Background(), REDIS_TIMEOUT_MILLI_TS*time.Millisecond)
	defer func() {
		ctxErr := ctx.Err()
		cancel()
		if err == nil {
			err = ctxErr
		}
	}()

	_, err = rdb.Del(ctx, "lock:"+key).Result()
	if err != nil {
		return err
	}

	return nil
}
