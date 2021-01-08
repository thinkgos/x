package czlib

import (
	"bytes"
	"compress/gzip"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/thinkgos/x/internal/mock"
)

func TestConn(t *testing.T) {
	t.Run("invalid zlib", func(t *testing.T) {
		data := []byte("aaaaa")

		mconn := mock.New(new(bytes.Buffer))
		conn := New(mconn)

		_, err := conn.Read(data)
		require.Error(t, err)
	})

	t.Run("invalid level", func(t *testing.T) {
		mconn := mock.New(new(bytes.Buffer))
		conn := NewLevel(mconn, gzip.HuffmanOnly-1)

		_, err := conn.Write([]byte("aaaaaa"))
		require.Error(t, err)
	})

	t.Run("zlib", func(t *testing.T) {
		data := []byte(
			`hello worldhello worldhello worldhello worldhello worldhello worldhello world
hello worldhello worldhello worldhello worldhello worldhello worldhello world
hello worldhello worldhello worldhello worldhello worldhello worldhello world
hello worldhello worldhello worldhello worldhello worldhello worldhello world
hello worldhello worldhello worldhello worldhello worldhello worldhello world
hello worldhello worldhello worldhello worldhello worldhello worldhello world
hello worldhello worldhello worldhello worldhello worldhello worldhello world`,
		)

		buf := new(bytes.Buffer)

		mconn := mock.New(buf)
		conn := New(mconn)

		start := time.Now()

		// write
		n, err := conn.Write(data)
		require.NoError(t, err)
		require.Equal(t, len(data), n)

		t.Log(time.Since(start).String(), buf.Len())

		// read
		rd := make([]byte, len(data))
		n, err = conn.Read(rd)
		require.NoError(t, err)
		require.Equal(t, len(data), n)

		// same
		require.Equal(t, rd[:n], data)
	})
}
