package pool

import (
	"io"
	"net/http"
	"sync"
	"time"
)

type Connection struct {
	client *http.Client
}

func (c *Connection) Get(url string) (*http.Response, error) {
	return c.client.Get(url)
}

func (c *Connection) Post(url, contentType string, body io.Reader) (resp *http.Response, err error) {
	return c.client.Post(url, contentType, body)
}

// ConnectionPool 连接池结构体
type ConnectionPool struct {
	pool sync.Pool
}

// NewConnectionPool 创建连接池
func NewConnectionPool() *ConnectionPool {
	return &ConnectionPool{
		pool: sync.Pool{
			New: func() interface{} {
				// 创建一个新的连接
				transport := &http.Transport{
					MaxIdleConns:        50,               // 最大空闲连接数
					MaxIdleConnsPerHost: 20,               // 每个主机的最大空闲连接数
					IdleConnTimeout:     30 * time.Second, // 空闲连接的超时时间
				}
				client := &http.Client{
					Transport: transport,
					Timeout:   60 * time.Second, // 请求超时时间
				}
				return &Connection{client: client}
			},
		},
	}
}

// GetConnection 从连接池中获取连接
func (p *ConnectionPool) GetConnection() *Connection {
	return p.pool.Get().(*Connection)
}

// PutConnection 将连接放回连接池
func (p *ConnectionPool) PutConnection(conn *Connection) {
	p.pool.Put(conn)
}
