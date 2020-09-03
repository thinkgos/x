package encrypt

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewRc4Md5(t *testing.T) {
	key := []byte("0123456789abcdef")
	iv := []byte("0123456789abcdef")

	_, err := NewRc4Md5([]byte("invalidkey"), []byte("invalidiv"))
	require.Contains(t, err.Error(), "key")

	_, err = NewRc4Md5(key, []byte("invalidiv"))
	require.Contains(t, err.Error(), "iv")

	st, err := NewRc4Md5(key, iv)
	require.NoError(t, err)
	require.NotNil(t, st)
}

func TestNewSalsa20(t *testing.T) {
	key := []byte("0123456789abcdef0123456789abcdef")
	iv := []byte("01234567")

	_, err := NewSalsa20([]byte("invalidkey"), []byte("invalidiv"))
	require.Contains(t, err.Error(), "key")

	_, err = NewSalsa20(key, []byte("invalidiv"))
	require.Contains(t, err.Error(), "iv")

	st, err := NewSalsa20(key, iv)
	require.NoError(t, err)
	require.NotNil(t, st)
}
