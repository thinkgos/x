package bpool

import (
	"sync"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSyncPool(t *testing.T) {
	p := NewPool(2048)
	b := p.Get()
	bs := b[0:cap(b)]
	require.Equal(t, cap(b), len(bs))
	p.Put(b)
	p.Get()
	p.Put(b)
	require.Panics(t, func() { p.Put([]byte{}) })
}

func BenchmarkSyncPool(b *testing.B) {
	p := NewPool(2048)
	wg := new(sync.WaitGroup)

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		wg.Add(1)
		go func() {
			s := p.Get()
			p.Put(s)
			wg.Done()
		}()
	}
	wg.Wait()
}
