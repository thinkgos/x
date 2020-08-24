package ciol

import (
	"bytes"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"golang.org/x/time/rate"

	"github.com/thinkgos/go-core-package/internal/mock"
)

func mockConn(t *testing.T, data []byte, r, d int, options ...Options) *Conn {
	mconn := mock.New(new(bytes.Buffer))

	conn := New(
		mconn,
		options...,
	)
	assert.Equal(t, rate.Limit(r), conn.ReadLimit())
	assert.Equal(t, rate.Limit(d), conn.WriteLimit())

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

	return conn
}

func TestConn_no_limiter(t *testing.T) {
	data := []byte("hello world")
	conn := mockConn(t, data, 0, 0)
	defer conn.Close() //nolint: errcheck

	conn.SetReadLimit(1024)
	assert.Equal(t, rate.Limit(0), conn.ReadLimit())
	conn.SetWriteLimit(1024)
	assert.Equal(t, rate.Limit(0), conn.WriteLimit())

	conn.SetReadLimitAt(time.Now(), 2048)
	assert.Equal(t, rate.Limit(0), conn.ReadLimit())
	conn.SetWriteLimitAt(time.Now(), 2048)
	assert.Equal(t, rate.Limit(0), conn.WriteLimit())

	conn.SetReadBurst(1024)
	conn.SetReadBurstAt(time.Now(), 1024)
	conn.SetWriteBurst(1024)
	conn.SetWriteBurstAt(time.Now(), 1024)
}

func TestConn_read_limiter(t *testing.T) {
	data := []byte("hello world")
	conn := mockConn(t, data, 100, 0,
		WithReadLimiter(rate.Limit(100), 10240),
	)
	conn.Close() //nolint: errcheck
}

func TestConn_write_limiter(t *testing.T) {
	data := []byte("hello world")
	conn := mockConn(t, data, 0, 100,
		WithWriteLimiter(rate.Limit(100), 10240),
	)
	conn.Close() //nolint: errcheck
}

func TestConn_read_write_limiter(t *testing.T) {
	data := []byte("hello world")
	conn := mockConn(t, data, 100, 100,
		WithWriteLimiter(rate.Limit(100), 10240),
		WithReadLimiter(rate.Limit(100), 10240),
	)
	defer conn.Close() //nolint: errcheck

	conn.SetReadLimit(1024)
	assert.Equal(t, rate.Limit(1024), conn.ReadLimit())
	conn.SetWriteLimit(1024)
	assert.Equal(t, rate.Limit(1024), conn.WriteLimit())

	conn.SetReadLimitAt(time.Now(), 2048)
	assert.Equal(t, rate.Limit(2048), conn.ReadLimit())
	conn.SetWriteLimitAt(time.Now(), 2048)
	assert.Equal(t, rate.Limit(2048), conn.WriteLimit())

	conn.SetReadBurst(1024)
	conn.SetReadBurstAt(time.Now(), 1024)
	conn.SetWriteBurst(1024)
	conn.SetWriteBurstAt(time.Now(), 1024)
}
