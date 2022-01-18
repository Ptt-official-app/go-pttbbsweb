package boardd

import (
	context "context"

	grpc "google.golang.org/grpc"
)

type MockClientConn struct{}

func NewMockClientConn() *MockClientConn {
	return &MockClientConn{}
}

func (c *MockClientConn) Invoke(ctx context.Context, method string, args interface{}, reply interface{}, opts ...grpc.CallOption) (err error) {
	switch method {
	case "/pttbbs.api.BoardService/Hotboard":
		r := reply.(*HotboardReply)
		r.Boards = []*Board{testBoard10, testBoard1, testBoard8}
	}
	return nil
}

func (c *MockClientConn) NewStream(ctx context.Context, desc *grpc.StreamDesc, method string, opts ...grpc.CallOption) (cs grpc.ClientStream, err error) {
	return nil, nil
}

func (c *MockClientConn) Close() {
}
