package queue

import (
	"github.com/golang-queue/queue"
	"github.com/golang-queue/queue/simple"
)

var client *queue.Queue

func Start() error {
	var err error
	// define the worker
	w := simple.NewWorker(
		simple.WithQueueNum(N_COMMENT_QUEUE),
	)

	// define the queue
	client, err = queue.NewQueue(
		queue.WithWorker(w),
	)
	if err != nil {
		return err
	}

	// start the worker
	client.Start()

	return nil
}

func Close() {
	// shutdown the service and notify all the worker
	client.Shutdown()
	// wait all jobs are complete.
	client.Wait()
}
