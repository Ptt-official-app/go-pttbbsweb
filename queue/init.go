package queue

import (
	"context"
	"encoding/json"

	"github.com/appleboy/queue"
	"github.com/appleboy/queue/simple"
)

var client *queue.Queue

func Start() error {
	var err error
	// define the worker
	w := simple.NewWorker(
		simple.WithQueueNum(N_COMMENT_QUEUE),
		simple.WithRunFunc(func(ctx context.Context, m queue.QueuedMessage) error {
			v, ok := m.(*CommentQueue)
			if !ok {
				if err := json.Unmarshal(m.Bytes(), &v); err != nil {
					return err
				}
			}
			return ProcessCommentQueue(v)
		}),
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
