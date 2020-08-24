package cbuffered

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/thinkgos/go-core-package/internal/mock"
)

func MockConn(t *testing.T, data []byte, sizes ...int) {
	mconn := mock.New(new(bytes.Buffer))
	conn := New(mconn, sizes...)

	// write
	n, err := conn.Write(data)
	require.NoError(t, err)
	require.Equal(t, len(data), n)

	pData, err := conn.Peek(len(data))
	require.NoError(t, err)
	require.Equal(t, data, pData)

	assert.Equal(t, len(data), conn.Buffered())

	b, err := conn.ReadByte()
	require.NoError(t, err)
	require.Equal(t, data[0], b)

	assert.Equal(t, len(data)-1, conn.Buffered())

	err = conn.UnreadByte()
	require.NoError(t, err)

	assert.Equal(t, len(data), conn.Buffered())

	// read
	rd := make([]byte, len(data))
	n, err = conn.Read(rd)
	require.NoError(t, err)
	require.Equal(t, len(data), n)

	assert.Equal(t, 0, conn.Buffered())

	// same
	require.Equal(t, rd[:n], data)
}

func TestConn(t *testing.T) {
	data := []byte("hello world")

	MockConn(t, data)
	MockConn(t, data, len(data))
}
