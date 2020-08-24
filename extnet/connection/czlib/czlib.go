package czlib

import (
	"compress/zlib"
	"net"
)

// Conn is a generic stream-oriented network connection with gzip
type Conn struct {
	net.Conn
	level int
	dict  []byte
}

// New new a zlib compress
func New(conn net.Conn) *Conn {
	return NewLevel(conn, zlib.DefaultCompression)
}

// NewLevel new a zlib compress with the level
// level see zlib package
func NewLevel(conn net.Conn, level int) *Conn {
	return NewLevelDict(conn, level, nil)
}

// NewLevelDict new a zlib compress with the level and dict
// level see zlib package
func NewLevelDict(conn net.Conn, level int, dict []byte) *Conn {
	return &Conn{
		conn,
		level,
		dict,
	}
}

// Read reads data from the connection.
func (sf *Conn) Read(p []byte) (int, error) {
	r, err := zlib.NewReaderDict(sf.Conn, sf.dict)
	if err != nil {
		return 0, err
	}
	return r.Read(p)
}

// Write writes data to the connection.
func (sf *Conn) Write(p []byte) (int, error) {
	w, err := zlib.NewWriterLevelDict(sf.Conn, sf.level, sf.dict)
	if err != nil {
		return 0, err
	}
	n, _ := w.Write(p)
	err = w.Flush()
	return n, err
}
