package queue

var N_COMMENT_QUEUE = 4096

func config() {
	N_COMMENT_QUEUE = setIntConfig("N_COMMENT_QUEUE", N_COMMENT_QUEUE)
}
