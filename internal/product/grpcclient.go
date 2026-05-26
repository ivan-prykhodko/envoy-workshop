package product

import (
	"sync"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/keepalive"
)

type Connection interface {
	ClientConn() (*grpc.ClientConn, error)
	Close() error
}

type connection struct {
	addr string
	conn *grpc.ClientConn
	mu   sync.Mutex
}

func NewGRpcConnection() Connection {
	return &connection{
		addr: "product-envoy:15001",
	}
}

func (c *connection) ClientConn() (*grpc.ClientConn, error) {
	c.mu.Lock()
	defer c.mu.Unlock()

	if c.conn != nil {
		return c.conn, nil
	}

	var err error
	c.conn, err = grpc.NewClient(
		c.addr,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithKeepaliveParams(
			keepalive.ClientParameters{
				Time:                60 * time.Second,
				Timeout:             5 * time.Second,
				PermitWithoutStream: true,
			},
		),
	)

	return c.conn, err
}

func (c *connection) Close() error {
	c.mu.Lock()
	defer c.mu.Unlock()

	if c.conn == nil {
		return nil
	}

	err := c.conn.Close()
	c.conn = nil

	return err
}
