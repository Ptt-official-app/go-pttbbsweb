package schema

import (
	"context"
	"time"

	"github.com/Ptt-official-app/go-pttbbs/bbs"
)

func Set2FA(userID bbs.UUserID, email string, token string, expireTSDuration time.Duration) (err error) {
	ctx, cancel := context.WithTimeout(context.Background(), REDIS_TIMEOUT_MILLI_TS*time.Millisecond)
	defer func() {
		ctxErr := ctx.Err()
		cancel()
		if err == nil {
			err = ctxErr
		}
	}()

	value := TwoFactorSerializeValue(email, token)
	val, err := rdb.SetNX(ctx, "2fa:"+string(userID), value, expireTSDuration).Result()
	if err != nil {
		return err
	}
	if !val {
		return Err2FAAlreadyExists
	}

	return nil
}

func TwoFactorSerializeValue(email string, token string) string {
	return token + ":" + email
}

func Get2FA(userID bbs.UUserID) (emailtoken string, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), REDIS_TIMEOUT_MILLI_TS*time.Millisecond)
	defer func() {
		ctxErr := ctx.Err()
		cancel()
		if err == nil {
			err = ctxErr
		}
	}()

	emailtoken, err = rdb.Get(ctx, "2fa:"+string(userID)).Result()
	if err != nil {
		return "", err
	}

	return emailtoken, nil
}
