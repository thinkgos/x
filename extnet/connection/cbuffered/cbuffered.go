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

// Package cbuffered 实现读缓冲的net.conn接口
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
