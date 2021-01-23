package schema

import (
	"context"
	"time"

	"github.com/Ptt-official-app/go-pttbbs/bbs"
)

func Set2FA(userID bbs.UUserID, token string, expireTSDuration time.Duration) (err error) {

	ctx, cancel := context.WithTimeout(context.Background(), REDIS_TIMEOUT_MILLI_TS*time.Millisecond)
	defer func() {
		ctxErr := ctx.Err()
		cancel()
		if err == nil {
			err = ctxErr
		}
	}()

	val, err := rdb.SetNX(ctx, "2fa:"+string(userID), token, expireTSDuration).Result()
	if err != nil {
		return err
	}
	if !val {
		return ErrNoLock
	}

	return nil

}

func Get2FA(userID bbs.UUserID) (token string, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), REDIS_TIMEOUT_MILLI_TS*time.Millisecond)
	defer func() {
		ctxErr := ctx.Err()
		cancel()
		if err == nil {
			err = ctxErr
		}
	}()

	token, err = rdb.Get(ctx, "2fa:"+string(userID)).Result()
	if err != nil {
		return "", err
	}

	return token, nil
}
