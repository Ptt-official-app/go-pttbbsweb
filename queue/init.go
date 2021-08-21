package queue

import (
	"runtime"

	"github.com/golang-queue/queue"
)

var client *queue.Queue

func Start() error {
	client = queue.NewPool(
		runtime.NumCPU(),
		queue.WithQueueSize(N_COMMENT_QUEUE),
	)

	// start the worker
	client.Start()

	return nil
}

func Close() {
	// shutdown the service and notify all the worker
	// wait all jobs are complete.
	client.Release()
}
