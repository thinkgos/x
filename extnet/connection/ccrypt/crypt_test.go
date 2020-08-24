package ccrypt

import (
	"bytes"
	"crypto/sha1"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/thinkgos/go-core-package/internal/mock"
)

func mockWithConfig(t *testing.T, data []byte, cfg Config) {
	mconn := mock.New(new(bytes.Buffer))
	conn := New(mconn, cfg)

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
}

func TestConn(t *testing.T) {
	data := []byte("hello world")

	mockWithConfig(t, data, Config{
		Password: "password",
	})
	mockWithConfig(t, data, Config{
		Password:   "password",
		Salt:       DefaultSalt,
		Iterations: DefaultIterations,
		KeySize:    16,
		Hash:       sha1.New,
	})
	mockWithConfig(t, data, Config{
		Password:   "password",
		Salt:       DefaultSalt,
		Iterations: DefaultIterations,
		KeySize:    24,
		Hash:       sha1.New,
	})
}
