package grpc_client

import (
	"ai-chat-backend/pkg/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/connectivity"
	"sync"
)

type ClientPool interface {
	Get() *grpc.ClientConn
	Put(*grpc.ClientConn)
}

type clientPool struct {
	pool sync.Pool
}

func NewPool(target string, opts ...grpc.DialOption) (ClientPool, error) {
	return &clientPool{
		pool: sync.Pool{
			New: func() any {
				conn, err := grpc.NewClient(target, opts...)
				if err != nil {
					log.Error(err)
					return nil
				}
				return conn
			},
		},
	}, nil
}

func (c *clientPool) Get() *grpc.ClientConn {
	conn := c.pool.Get().(*grpc.ClientConn)
	if conn.GetState() == connectivity.Shutdown || conn.GetState() == connectivity.TransientFailure {
		conn.Close()
		conn = c.pool.New().(*grpc.ClientConn)
	}
	return conn
}
func (c *clientPool) Put(conn *grpc.ClientConn) {
	if conn.GetState() == connectivity.Shutdown || conn.GetState() == connectivity.TransientFailure {
		conn.Close()
		return
	}
	c.pool.Put(conn)
}
