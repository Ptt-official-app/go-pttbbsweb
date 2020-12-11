package backend

const (
	LOGIN               = "/token"
	REGISTER            = "/register"
	LOAD_GENERAL_BOARDS = "/board/boards"
)

func WithPrefix(route string) string {
	return HTTP_PREFIX + route
}
