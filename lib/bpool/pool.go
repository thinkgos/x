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

package bpool

import (
	"sync"
)

// Pool pool buffer
type Pool struct {
	size int
	pool *sync.Pool
}

// NewPool creates a leaky buffer which can hold at most n buffer, each
// with capacity bytes.
func NewPool(capacity int) *Pool {
	return &Pool{
		capacity,
		&sync.Pool{
			New: func() interface{} {
				return make([]byte, 0, capacity)
			}},
	}
}

// Get selects an arbitrary item from the Pool, removes it from the
// Pool, and returns it to the caller.
func (sf *Pool) Get() []byte {
	return sf.pool.Get().([]byte)
}

// Put adds x to the pool.
func (sf *Pool) Put(b []byte) {
	if cap(b) != sf.size {
		panic("invalid buffer capacity that's put into leaky buffer")
	}
	sf.pool.Put(b[:0]) // nolint: staticcheck
}
