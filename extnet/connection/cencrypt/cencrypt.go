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

// Package cencrypt 实现加密的net.conn接口
package cencrypt

import (
	"crypto/cipher"
	"io"
	"net"

	"github.com/thinkgos/x/lib/encrypt"
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
