package cflow

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/atomic"

	"github.com/thinkgos/go-core-package/internal/mock"
)

func TestConn(t *testing.T) {
	data := []byte("hello world")

	mconn := mock.New(new(bytes.Buffer))
	conn := Conn{
		mconn,
		atomic.NewUint64(0),
		atomic.NewUint64(0),
		atomic.NewUint64(0),
	}

	// write
	n, err := conn.Write(data)
	require.NoError(t, err)
	require.Equal(t, len(data), n)

	// read
	rd := make([]byte, len(data))
	n, err = conn.Read(rd)
	require.NoError(t, err)
	require.Equal(t, len(data), n)

	// same
	require.Equal(t, rd[:n], data)

	assert.Equal(t, uint64(len(data)), conn.Wc.Load())
	assert.Equal(t, uint64(len(data)), conn.Rc.Load())
	assert.Equal(t, uint64(len(data))*2, conn.Tc.Load())
}
