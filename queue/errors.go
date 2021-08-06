package queue

import (
	"fmt"
	"time"
)

const defaultTimeout = 100 * time.Millisecond

type ErrTimeout struct {
	timeout time.Duration
}

func (e ErrTimeout) Error() string {
	return fmt.Sprintf("timeout: %s", e.timeout)
}
