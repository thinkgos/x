// Package csnappy 采用snappy压缩实现的net.conn接口
package csnappy

import (
	"net"

	"github.com/golang/snappy"
)

// Conn is a generic stream-oriented network connection with snappy
type Conn struct {
	net.Conn
	w *snappy.Writer
	r *snappy.Reader
}

// New new with snappy
func New(conn net.Conn) *Conn {
	return &Conn{
		conn,
		snappy.NewBufferedWriter(conn),
		snappy.NewReader(conn),
	}
}

// Read reads data from the connection.
func (sf *Conn) Read(p []byte) (int, error) {
	return sf.r.Read(p)
}

// Write writes data to the connection.
func (sf *Conn) Write(p []byte) (int, error) {
	n, _ := sf.w.Write(p)
	err := sf.w.Flush()
	return n, err
}
