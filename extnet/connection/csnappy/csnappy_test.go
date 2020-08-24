package csnappy

import (
	"bytes"
	"testing"
	"time"

	"github.com/golang/snappy"
	"github.com/stretchr/testify/require"

	"github.com/thinkgos/go-core-package/internal/mock"
)

func TestConn(t *testing.T) {
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

	t.Log(time.Since(start).String(), len(data), buf.Len())

	// read
	rd := make([]byte, len(data))
	n, err = conn.Read(rd)
	require.NoError(t, err)
	require.Equal(t, len(data), n)

	// same
	require.Equal(t, rd[:n], data)
}

func BenchmarkSnappy(b *testing.B) {
	var testData = []byte(`hello worldhello worldhello worldhello worldhello worldhello worldhello world
hello worldhello worldhello worldhello worldhello worldhello worldhello world
hello worldhello worldhello worldhello worldhello worldhello worldhello world
hello worldhello worldhello worldhello worldhello worldhello worldhello world
hello worldhello worldhello worldhello worldhello worldhello worldhello world
hello worldhello worldhello worldhello worldhello worldhello worldhello world
hello worldhello worldhello worldhello worldhello worldhello worldhello world`,
	)

	dst := make([]byte, len(testData))
	for i := 0; i < b.N; i++ {
		snappy.Encode(dst, testData)
	}
}
