package extnet

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/atomic"

	"github.com/thinkgos/go-core-package/internal/mock"
)

func TestAdorn(t *testing.T) {
	want := []byte(`
this is a testing mock!this is a testing mock!this is a testing mock!this is a testing mock!this is a testing mock!
this is a testing mock!this is a testing mock!this is a testing mock!this is a testing mock!this is a testing mock!
this is a testing mock!this is a testing mock!this is a testing mock!this is a testing mock!this is a testing mock!
this is a testing mock!this is a testing mock!this is a testing mock!this is a testing mock!this is a testing mock!
this is a testing mock!this is a testing mock!this is a testing mock!this is a testing mock!this is a testing mock!
this is a testing mock!this is a testing mock!this is a testing mock!this is a testing mock!this is a testing mock!
this is a testing mock!this is a testing mock!this is a testing mock!this is a testing mock!this is a testing mock!
this is a testing mock!this is a testing mock!this is a testing mock!this is a testing mock!this is a testing mock!
`)
	for _, compress := range []bool{true, false} {
		wc := atomic.NewUint64(0)
		rc := atomic.NewUint64(0)
		tc := atomic.NewUint64(0)

		chains := AdornConnsChain{
			AdornIol(),
			AdornFlow(wc, rc, tc),
			AdornSnappy(compress),
			AdornGzip(!compress),
			AdornZlib(compress),
		}

		buf := new(bytes.Buffer)
		conn := mock.New(buf)

		for _, chain := range chains {
			conn = chain(conn)
		}

		nw, err := conn.Write(want)
		require.NoError(t, err)
		require.Equal(t, len(want), nw)

		nb := buf.Len()

		got := make([]byte, 1024)
		nr, err := conn.Read(got)
		require.NoError(t, err)
		require.Equal(t, len(want), nr)
		require.Equal(t, want, got[:nr])

		assert.Equal(t, uint64(nb), wc.Load())
		assert.Equal(t, uint64(nb), rc.Load())
		assert.Equal(t, wc.Load()+rc.Load(), tc.Load())
	}

}
