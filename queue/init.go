package queue

func Init() {
	if theQueue != nil {
		return
	}
	theQueue = make(chan *CommentQueue, N_COMMENT_QUEUE)
	theQuit = make([]chan struct{}, N_COMMENT_QUEUE)

	for idx := 0; idx < N_COMMENT_QUEUE; idx++ {
		theQuit[idx] = make(chan struct{})
		go ProcessCommentQueue(idx, theQuit[idx])
	}
}

func Close() {
	for idx := 0; idx < N_COMMENT_QUEUE; idx++ {
		theQuit[idx] <- struct{}{}
	}

	close(theQueue)

	theQueue = nil
	theQuit = nil
}
