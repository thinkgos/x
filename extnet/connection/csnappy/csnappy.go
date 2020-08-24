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
