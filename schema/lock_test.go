package schema

import (
	"testing"
	"time"
)

func TestTryLock(t *testing.T) {
	setupTest()
	defer teardownTest()

	expireNanoTS := time.Duration(2) * time.Second
	type args struct {
		key    string
		expire time.Duration
	}
	tests := []struct {
		name        string
		args        args
		sleepTS     int
		isUnlock    bool
		wantErr     bool
		expectedErr error
	}{
		// TODO: Add test cases.
		{
			args: args{key: "test1", expire: expireNanoTS},
		},
		{
			args:        args{key: "test1", expire: expireNanoTS},
			wantErr:     true,
			expectedErr: ErrNoLock,
		},
		{
			args:     args{key: "test1", expire: expireNanoTS},
			isUnlock: true,
		},
		{
			args: args{key: "test2", expire: expireNanoTS},
		},
		{
			args:    args{key: "test1", expire: expireNanoTS},
			sleepTS: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.sleepTS > 0 {
				time.Sleep(time.Duration(tt.sleepTS) * time.Second)
			}
			if tt.isUnlock {
				Unlock(tt.args.key)
			}
			err := TryLock(tt.args.key, tt.args.expire)

			if (err != nil) != tt.wantErr {
				t.Errorf("TryLock() error = %v, wantErr %v", err, tt.wantErr)
			}

			if err != tt.expectedErr {
				t.Errorf("TryLock: e: %v expected: %v", err, tt.expectedErr)
			}
		})
	}
}
