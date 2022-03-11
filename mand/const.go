package mand

import (
	"sync"

	"google.golang.org/grpc"
)

var Cli ManServiceClient

var (
	conn    *grpc.ClientConn
	cliLock sync.Mutex
)
