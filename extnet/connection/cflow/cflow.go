// Package cflow 实现字节统计,读,写,读写统计,以字节为准. 三个参数为空时,无任何统计
package cflow

import (
	"net"

	"go.uber.org/atomic"
)

// Conn conn with read,write,total count
type Conn struct {
	net.Conn
	Wc *atomic.Uint64 // 写统计
	Rc *atomic.Uint64 // 读统计
	Tc *atomic.Uint64 // 读写统计
}

// Read reads data from the connection.
func (sf *Conn) Read(p []byte) (int, error) {
	n, err := sf.Conn.Read(p)
	if n != 0 {
		cnt := uint64(n)
		if sf.Rc != nil {
			sf.Rc.Add(cnt)
		}
		if sf.Tc != nil {
			sf.Tc.Add(cnt)
		}
	}
	return n, err
}

// Write writes data to the connection.
func (sf *Conn) Write(p []byte) (int, error) {
	n, err := sf.Conn.Write(p)
	if n != 0 {
		cnt := uint64(n)
		if sf.Wc != nil {
			sf.Wc.Add(cnt)
		}
		if sf.Tc != nil {
			sf.Tc.Add(cnt)
		}
	}
	return n, err
}
