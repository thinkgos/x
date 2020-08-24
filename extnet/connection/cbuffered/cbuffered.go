package cbuffered

import (
	"bufio"
	"net"
)

// Conn conn with bufio.Reader
type Conn struct {
	net.Conn
	reader *bufio.Reader
}

// New with net.Conn and a new Reader whose buffer has at least the specified size
func New(c net.Conn, sizes ...int) *Conn {
	if len(sizes) > 0 {
		return &Conn{c, bufio.NewReaderSize(c, sizes[0])}
	}
	return &Conn{c, bufio.NewReader(c)}
}

// Peek returns the next n bytes without advancing the reader. The bytes stop
// being valid at the next read call. If Peek returns fewer than n bytes, it
// also returns an error explaining why the read is short. The error is
// ErrBufferFull if n is larger than b's buffer size.
func (sf *Conn) Peek(n int) ([]byte, error) {
	return sf.reader.Peek(n)
}

// Read reads data into p.
func (sf *Conn) Read(p []byte) (int, error) {
	return sf.reader.Read(p)
}

// ReadByte reads and returns a single byte.
func (sf *Conn) ReadByte() (byte, error) {
	return sf.reader.ReadByte()
}

// UnreadByte unreads the last byte. Only the most recently read byte can be unread.
func (sf *Conn) UnreadByte() error {
	return sf.reader.UnreadByte()
}

// Buffered returns the number of bytes that can be read from the current buffer.
func (sf *Conn) Buffered() int {
	return sf.reader.Buffered()
}
