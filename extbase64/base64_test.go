package outil

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestBase64(t *testing.T) {
	b := []byte{1, 2, 3, 4, 5, 6}
	Encode(b)

	orig := "helloworld"
	bs64 := EncodeString(orig)

	rawByte, err := Decode(bs64)
	require.NoError(t, err)
	require.Equal(t, orig, string(rawByte))

	raw, err := DecodeString(bs64)
	require.NoError(t, err)
	require.Equal(t, orig, raw)
}
