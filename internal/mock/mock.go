// Package mock simulate a net.Conn with io.ReadWriter
package mock

import (
	"io"
	"net"
	"time"
)

// Mock a simulate a net.Conn with io.ReadWriter
type Mock struct {
	rw io.ReadWriter
}

// New new with io.ReadWriter
func New(rw io.ReadWriter) net.Conn { return &Mock{rw} }

// Read reads up to len(p) bytes into p.
func (m *Mock) Read(b []byte) (n int, err error) { return m.rw.Read(b) }

// Write writes len(p) bytes from p to the underlying data stream.
func (m *Mock) Write(b []byte) (n int, err error) { return m.rw.Write(b) }

// Close closes the connection.
func (m *Mock) Close() error { return nil }

// LocalAddr returns the local network address.
func (m *Mock) LocalAddr() net.Addr { return nil }

// RemoteAddr returns the remote network address.
func (m *Mock) RemoteAddr() net.Addr { return nil }

// SetDeadline sets the read and write deadlines associated
// with the connection. It is equivalent to calling both
// SetReadDeadline and SetWriteDeadline.
func (m *Mock) SetDeadline(time.Time) error { return nil }

// SetReadDeadline sets the deadline for future Read calls
// and any currently-blocked Read call.
// A zero value for t means Read will not time out.
func (m *Mock) SetReadDeadline(time.Time) error { return nil }

// SetWriteDeadline sets the deadline for future Write calls
// and any currently-blocked Write call.
// Even if write times out, it may return n > 0, indicating that
// some of the data was successfully written.
// A zero value for t means Write will not time out.
func (m *Mock) SetWriteDeadline(time.Time) error { return nil }
