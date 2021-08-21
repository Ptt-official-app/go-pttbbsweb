package queue

import "runtime"

var (
	queueSize = 4096
	workerNum = runtime.NumCPU()
)

func config() {
	queueSize = setIntConfig("QUEUE_SIZE", queueSize)
	workerNum = setIntConfig("WORKER_NUM", workerNum)
}
