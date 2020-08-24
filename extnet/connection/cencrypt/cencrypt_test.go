package cencrypt

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/thinkgos/go-core-package/internal/mock"
	"github.com/thinkgos/go-core-package/lib/encrypt"
)

func TestConn(t *testing.T) {
	password := "password"
	data := []byte("hello world")

	mconn := mock.New(new(bytes.Buffer))

	for _, method := range encrypt.CipherMethods() {
		cip, err := encrypt.NewCipher(method, password)
		require.NoError(t, err)
		conn := New(mconn, cip)

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
}
