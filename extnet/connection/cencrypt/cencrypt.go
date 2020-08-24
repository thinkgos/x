package cencrypt

import (
	"crypto/cipher"
	"io"
	"net"

	"github.com/thinkgos/go-core-package/lib/encrypt"
)

// Conn conn with encrypt.Cipher
type Conn struct {
	net.Conn
	w io.Writer
	r io.Reader
}

// New a connection with encrypt cipher
func New(c net.Conn, cip *encrypt.Cipher) *Conn {
	return &Conn{
		c,
		&cipher.StreamWriter{S: cip.Write, W: c},
		&cipher.StreamReader{S: cip.Read, R: c},
	}
}

// Read reads data from the connection.
func (sf *Conn) Read(b []byte) (n int, err error) {
	return sf.r.Read(b)
}

// Write writes data to the connection.
func (sf *Conn) Write(b []byte) (n int, err error) {
	return sf.w.Write(b)
}
