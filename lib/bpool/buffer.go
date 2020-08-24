package bpool

// Buffer leaky buffer
type Buffer struct {
	size     int // size of each buffer
	freeList chan []byte
}

// NewBuffer creates a leaky buffer which can hold at most n buffer, each
// with size bytes.
func NewBuffer(maxFreeSize, size int) *Buffer {
	return &Buffer{
		size,
		make(chan []byte, maxFreeSize),
	}
}

// Get returns a buffer from the leaky buffer or create a new buffer.
func (sf *Buffer) Get() (b []byte) {
	select {
	case b = <-sf.freeList:
	default:
		b = make([]byte, 0, sf.size)
	}
	return
}

// Put add the buffer into the free buffer pool for reuse. Panic if the buffer
// size is not the same with the leaky buffer's. This is intended to expose
// error usage of leaky buffer.
func (sf *Buffer) Put(b []byte) {
	if cap(b) != sf.size {
		panic("invalid buffer size that's put into leaky buffer")
	}
	select {
	case sf.freeList <- b[:0]:
	default:
	}
}
