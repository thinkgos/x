package encrypt

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestEncrypt(t *testing.T) {
	_, err := NewCipher("invalid_method", "")
	require.Error(t, err)
	_, err = NewCipher("invalid_method", "pass_word")
	require.Error(t, err)

	password := "pass_word"
	src := []byte("hello world")
	for _, method := range CipherMethods() {
		require.True(t, HasCipherMethod(method))

		cip, err := NewCipher(method, password)
		require.NoError(t, err)

		// encrypt
		encVal := make([]byte, len(src))
		cip.Write.XORKeyStream(encVal, src)
		// decrypt
		decVal := make([]byte, len(encVal))
		cip.Read.XORKeyStream(decVal, encVal)

		require.Equal(t, decVal, src)
	}
}

func TestStream(t *testing.T) {
	_, err := NewStream("invalid_method", []byte{}, []byte{}, true)
	require.Error(t, err)
	_, err = NewStream("aes-128-cfb", []byte{}, []byte("01234567890abcdef"), true)
	require.Error(t, err)
	_, err = NewStream("aes-128-cfb", []byte("01234567890abcdef"), []byte{}, true)
	require.Error(t, err)

	password := "pass_word"
	key := Evp2Key(password, 32)
	iv := Evp2Key("password", 32)
	src := []byte("hello world")
	for _, method := range CipherMethods() {
		require.True(t, HasCipherMethod(method))

		wr, err := NewStream(method, key, iv, true)
		require.NoError(t, err)
		rd, err := NewStream(method, key, iv, false)
		require.NoError(t, err)

		// encrypt
		encVal := make([]byte, len(src))
		wr.XORKeyStream(encVal, src)
		// decrypt
		decVal := make([]byte, len(encVal))
		rd.XORKeyStream(decVal, encVal)

		require.Equal(t, decVal, src)
	}
}
