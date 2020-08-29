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

// Buffer leaky buffer
type Buffer struct {
	capacity int // capacity of each buffer
	freeList chan []byte
}

// NewBuffer creates a leaky buffer which can hold at most n buffer, each
// with capacity bytes.
func NewBuffer(maxFreeSize, capacity int) *Buffer {
	return &Buffer{
		capacity,
		make(chan []byte, maxFreeSize),
	}
}

// Get returns a buffer from the leaky buffer or create a new buffer.
func (sf *Buffer) Get() (b []byte) {
	select {
	case b = <-sf.freeList:
	default:
		b = make([]byte, 0, sf.capacity)
	}
	return
}

// Put add the buffer into the free buffer pool for reuse. Panic if the buffer
// capacity is not the same with the leaky buffer's. This is intended to expose
// error usage of leaky buffer.
func (sf *Buffer) Put(b []byte) {
	if cap(b) != sf.capacity {
		panic("invalid buffer capacity that's put into leaky buffer")
	}
	select {
	case sf.freeList <- b[:0]:
	default:
	}
}
