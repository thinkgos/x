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

// Package flow 实现字节统计,读,写,读写统计,以字节为准. 三个参数为空时,无任何统计
package flow

import (
	"io"

	"go.uber.org/atomic"
)

// Flow io流统计
type Flow struct {
	io.ReadWriter
	Wc *atomic.Uint64 // 写统计
	Rc *atomic.Uint64 // 读统计
	Tc *atomic.Uint64 // 读写统计
}

// Read reads up to len(p) bytes into p.
func (sf *Flow) Read(p []byte) (int, error) {
	n, err := sf.ReadWriter.Read(p)
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

// Write writes len(p) bytes from p to the underlying data stream.
func (sf *Flow) Write(p []byte) (int, error) {
	n, err := sf.ReadWriter.Write(p)
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
