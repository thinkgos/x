package flow

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/require"
	"go.uber.org/atomic"
)

func TestFlow(t *testing.T) {
	var wc atomic.Uint64
	var rc atomic.Uint64
	var tc atomic.Uint64

	c := Flow{
		new(bytes.Buffer),
		&wc,
		&rc,
		&tc,
	}

	count := uint64(10240)
	c.Write(make([]byte, count)) // nolint: errcheck
	c.Read(make([]byte, count))  // nolint: errcheck

	require.Equal(t, count, wc.Load())
	require.Equal(t, count, rc.Load())
	require.Equal(t, count*2, tc.Load())
}
