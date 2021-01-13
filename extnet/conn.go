// Copyright [2020] [thinkgos] thinkgo@aliyun.com
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package extnet

import (
	"context"
	"net"
	"time"
)

// Dialer A Dialer is a means to establish a connection.
type Dialer interface {
	Dial(network, address string) (net.Conn, error)
}

// ContextDialer A ContextDialer dials using a context.
type ContextDialer interface {
	DialContext(ctx context.Context, network, address string) (net.Conn, error)
}

// AdornConn defines the conn decorate.
type AdornConn func(conn net.Conn) net.Conn

// AdornConnsChain defines a adornConn array.
// NOTE: 在conn read或write调用过程是在链上从后往前执行的,(类似洋葱,包在最外面的选执行),
//  所以基础链,统计链的应放在链头,也就是chains的第一个,最靠近出口
type AdornConnsChain []AdornConn

// Client tcp dialer
type Client struct {
	Timeout     time.Duration   // timeout for dial
	AdornChains AdornConnsChain // adorn chains
	Forward     Dialer          // if set it will use forward.
}

// Dial connects to the address on the named network.
func (sf *Client) Dial(network, addr string) (net.Conn, error) {
	return sf.DialContext(context.Background(), network, addr)
}

// DialContext connects to the address on the named network using the provided context.
func (sf *Client) DialContext(ctx context.Context, network, addr string) (net.Conn, error) {
	var d Dialer = &net.Dialer{Timeout: sf.Timeout}

	if sf.Forward != nil {
		d = sf.Forward
	}

	contextDial := func(ctx context.Context, network, addr string) (net.Conn, error) {
		return DialContext(ctx, d, network, addr)
	}
	if f, ok := d.(ContextDialer); ok {
		contextDial = f.DialContext
	}

	conn, err := contextDial(ctx, network, addr)
	if err != nil {
		return nil, err
	}
	for _, chain := range sf.AdornChains {
		conn = chain(conn)
	}
	return conn, nil
}

// DialContext dial context with dialer
// WARNING: this can leak a goroutine for as long as the underlying Dialer implementation takes to timeout
// A Conn returned from a successful Dial after the context has been canceled will be immediately closed.
func DialContext(ctx context.Context, d Dialer, network, address string) (net.Conn, error) {
	var conn net.Conn
	var err error

	done := make(chan struct{}, 1)
	go func() {
		conn, err = d.Dial(network, address)
		close(done)
		if conn != nil && ctx.Err() != nil {
			conn.Close()
		}
	}()
	select {
	case <-ctx.Done():
		err = ctx.Err()
	case <-done:
	}
	return conn, err
}

type listener struct {
	net.Listener
	AdornChains AdornConnsChain
}

// Listen announces on the local network address and afterChains
func Listen(network, addr string, chains ...AdornConn) (net.Listener, error) {
	l, err := net.Listen(network, addr)
	if err != nil {
		return nil, err
	}
	return NewListener(l, chains...), nil
}

// NewListener new listener
func NewListener(inner net.Listener, chains ...AdornConn) net.Listener {
	l := new(listener)
	l.Listener = inner
	l.AdornChains = chains
	return l
}

// Accept waits for and returns the next incoming TLS connection.
// The returned connection is of type *Conn.
func (l *listener) Accept() (net.Conn, error) {
	c, err := l.Listener.Accept()
	if err != nil {
		return nil, err
	}
	for _, chain := range l.AdornChains {
		c = chain(c)
	}
	return c, nil
}
