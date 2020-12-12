package backend

func WithPrefix(route string) string {
	return HTTP_PREFIX + route
}
