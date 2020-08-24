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
