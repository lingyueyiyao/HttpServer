package http

import (
	"golang.org/x/net/context"
	"net"
	"time"
)

// MyDialer struct 初始化最佳实践
type MyDialer interface {
	DialContext(context.Context, net.Addr) (net.Conn, error)
}

type myDialer struct {
	dialer *net.Dialer
	timeout time.Duration
	retry int
	MyDialer
}

func NewMyDialer(opts ...Option) MyDialer {
	// 1. 初始化默认值
	d := &myDialer{
		timeout: time.Second,
		retry: 5,
	}

	// 2. 加载option
	for _, opt := range opts {
		opt.apply(d)
	}

	// 3. 检查特殊值
	if d.dialer == nil {
		d.dialer = &net.Dialer{}
	}

	return d
}
