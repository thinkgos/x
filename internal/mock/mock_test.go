package mock

import (
	"bytes"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMock(t *testing.T) {
	data := []byte("hello world")

	buf := new(bytes.Buffer)
	mk := New(buf)

	n, err := mk.Write(data)
	require.NoError(t, err)
	assert.Equal(t, len(data), n)

	// check buf
	assert.Equal(t, len(data), buf.Len())
	require.Equal(t, data, buf.Bytes())

	got := make([]byte, len(data))
	n, err = mk.Read(got)
	require.NoError(t, err)
	assert.Equal(t, len(data), n)
	assert.Equal(t, data, got)

	// check buf
	assert.Equal(t, 0, buf.Len())

	// improve coverage
	mk.LocalAddr()
	mk.RemoteAddr()
	mk.SetDeadline(time.Now())      // nolint: errcheck
	mk.SetReadDeadline(time.Now())  // nolint: errcheck
	mk.SetWriteDeadline(time.Now()) // nolint: errcheck
	mk.Close()                      // nolint: errcheck
}
