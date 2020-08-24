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
// with size bytes.
func NewPool(size int) *Pool {
	return &Pool{
		size,
		&sync.Pool{
			New: func() interface{} {
				return make([]byte, 0, size)
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
		panic("invalid buffer size that's put into leaky buffer")
	}
	sf.pool.Put(b[:0]) // nolint: staticcheck
}
