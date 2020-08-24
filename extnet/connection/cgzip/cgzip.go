// Package cgzip 采用gzip压缩实现的net.conn接口
package cgzip

import (
	"compress/gzip"
	"net"
)

// Conn is a generic stream-oriented network connection with gzip
type Conn struct {
	net.Conn
	level int
}

// New new a gzip compress
func New(conn net.Conn) *Conn {
	return NewLevel(conn, gzip.DefaultCompression)
}

// NewLevel new a gzip compress with the level.
// level see gzip package
func NewLevel(conn net.Conn, level int) *Conn {
	return &Conn{
		conn,
		level,
	}
}

// Read reads data from the connection.
func (sf *Conn) Read(p []byte) (int, error) {
	r, err := gzip.NewReader(sf.Conn)
	if err != nil {
		return 0, err
	}
	return r.Read(p)
}

// Write writes data to the connection.
func (sf *Conn) Write(p []byte) (int, error) {
	w, err := gzip.NewWriterLevel(sf.Conn, sf.level)
	if err != nil {
		return 0, err
	}
	n, _ := w.Write(p)
	err = w.Flush()
	return n, err
}
