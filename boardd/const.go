package boardd

import (
	"sync"

	grpc "google.golang.org/grpc"
)

// Cli: this will be used in api.
var Cli BoardServiceClient

var (
	conn *grpc.ClientConn

	cliLock sync.Mutex
)
