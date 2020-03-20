package pool

import (
	"go-lib/log"
	"sync"

	"google.golang.org/grpc"
)

var DefaultGrpcConnPool *GrpcConnPool

func init() {
	DefaultGrpcConnPool = NewGrpcConnPool()
}

type GrpcConnPool struct {
	sync.RWMutex
	conn map[string]*grpc.ClientConn
}

func NewGrpcConnPool() *GrpcConnPool {
	return &GrpcConnPool{
		conn: make(map[string]*grpc.ClientConn, 10),
	}
}

func (p *GrpcConnPool) AddNode(addr string) {
	p.Lock()
	defer p.Unlock()
	if conn, ok := p.conn[addr]; ok {
		conn.Close()
		delete(p.conn, addr)
	}
	var conn, err = grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		log.Error(err)
		return
	}
	p.conn[addr] = conn
}

func (p *GrpcConnPool) DelNode(addr string) {
	p.Lock()
	defer p.Unlock()
	if conn, ok := p.conn[addr]; ok {
		conn.Close()
		delete(p.conn, addr)
	}
}

func (p *GrpcConnPool) GetConn(addr string) (*grpc.ClientConn, bool) {
	p.RLock()
	defer p.RUnlock()
	conn, ok := p.conn[addr]
	return conn, ok
}

func (p *GrpcConnPool) GetConns() map[string]*grpc.ClientConn {
	p.RLock()
	defer p.RUnlock()
	return p.conn
}
func (p *GrpcConnPool) GetRandomConn() (string, *grpc.ClientConn) {
	p.RLock()
	defer p.RUnlock()
	for addr, conn := range p.conn {
		return addr, conn
	}
	return "", nil
}
