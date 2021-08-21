package queue

import (
	"runtime"

	"github.com/Ptt-official-app/go-openbbsmiddleware/configutil"

	"github.com/golang-queue/queue"
)

const configPrefix = "go-openbbsmiddleware:queue"

var client *queue.Queue

func Start() error {
	queueSize := configutil.SetIntConfig(configPrefix, "QUEUE_SIZE", 4096)
	workerNum := configutil.SetIntConfig(configPrefix, "WORKER_NUM", runtime.NumCPU())

	client = queue.NewPool(
		workerNum,
		queue.WithQueueSize(queueSize),
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
