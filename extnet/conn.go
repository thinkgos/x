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
// NOTE: 在conn read或write调用过程是在链上从后往前执行的,(类似栈,先进后执行,后进先执行),
//  所以统计类的应放在链头,也就是AfterChains的第一个,最靠近出口
type AdornConnsChain []AdornConn

// Client tcp dialer
type Client struct {
	Timeout          time.Duration   // timeout for dial
	BaseAdorn        AdornConn       // base adorn conn
	AfterAdornChains AdornConnsChain // chains after base
	Forward          Dialer          // if set it will use forward.
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
	if sf.BaseAdorn != nil {
		conn = sf.BaseAdorn(conn)
	}
	for _, chain := range sf.AfterAdornChains {
		conn = chain(conn)
	}
	return conn, nil
}

// DialContext dial context with dialer
// WARNING: this can leak a goroutine for as long as the underlying Dialer implementation takes to timeout
// A Conn returned from a successful Dial after the context has been cancelled will be immediately closed.
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
	BaseAdornConn AdornConn
	AfterChains   AdornConnsChain
}

// Listen announces on the local network address and afterChains
func Listen(network string, addr string, afterChains ...AdornConn) (net.Listener, error) {
	return ListenWith(network, addr, nil, afterChains...)
}

// ListenWith announces on the local network address , base  afterChains
func ListenWith(network string, addr string, base AdornConn, afterChains ...AdornConn) (net.Listener, error) {
	l, err := net.Listen(network, addr)
	if err != nil {
		return nil, err
	}
	return NewListener(l, base, afterChains...), nil
}

// NewListener new listener
func NewListener(inner net.Listener, base AdornConn, afterChains ...AdornConn) net.Listener {
	l := new(listener)
	l.Listener = inner
	l.BaseAdornConn = base
	l.AfterChains = afterChains
	return l
}

// Accept waits for and returns the next incoming TLS connection.
// The returned connection is of type *Conn.
func (l *listener) Accept() (net.Conn, error) {
	c, err := l.Listener.Accept()
	if err != nil {
		return nil, err
	}
	if l.BaseAdornConn != nil {
		c = l.BaseAdornConn(c)
	}
	for _, chain := range l.AfterChains {
		c = chain(c)
	}
	return c, nil
}
