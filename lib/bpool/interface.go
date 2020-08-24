// Package bpool Provides bpool buffer pool
package bpool

// BufferPool buffer pool interface
type BufferPool interface {
	Get() []byte
	Put([]byte)
}
