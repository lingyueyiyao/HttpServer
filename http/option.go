package http

import (
	"net"
	"time"
)

type Option interface {
	apply(*myDialer)
}

// 接口型函数 适用于单方法接口
type optFunc func(*myDialer)

func (f optFunc) apply(d *myDialer) {
	f(d)
}

func RetryOption(r int) Option {
	return optFunc(func(d *myDialer) {
		d.retry = r
	})
}

func TimeoutOption(timeout time.Duration) Option {
	return optFunc(func(d *myDialer) {
		d.timeout = timeout
	})
}

func DialerOption(dialer *net.Dialer) Option {
	return optFunc(func(d *myDialer) {
		d.dialer = dialer
	})
}

