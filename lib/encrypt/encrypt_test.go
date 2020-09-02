package encrypt

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestEncrypt(t *testing.T) {
	t.Run("invalid", func(t *testing.T) {
		_, err := NewCipher("invalid_method", "")
		require.Error(t, err)
		_, err = NewCipher("invalid_method", "pass_word")
		require.Error(t, err)
	})
	t.Run("key iv length", func(t *testing.T) {
		kv, ok := GetCipher("rc4-md5")
		require.True(t, ok)
		require.Equal(t, 16, kv.KeyLen())
		require.Equal(t, 16, kv.IvLen())

		kv, ok = GetCipher("aes-128-cfb")
		require.True(t, ok)
		require.Equal(t, 16, kv.KeyLen())
		require.Equal(t, 16, kv.IvLen())
	})
	t.Run("cipher", func(t *testing.T) {
		password := "pass_word"
		plainText := []byte("hello world")
		for _, method := range CipherMethods() {
			require.True(t, Valid(method, password))
			require.True(t, HasCipherMethod(method))

			cip, err := NewCipher(method, password)
			require.NoError(t, err)

			// encrypt
			cipherText := make([]byte, len(plainText))
			cip.Write.XORKeyStream(cipherText, plainText)
			// decrypt
			got := make([]byte, len(cipherText))
			cip.Read.XORKeyStream(got, cipherText)

			require.Equal(t, plainText, got, fmt.Errorf("method: %s", method))
		}
	})
}

func TestStream(t *testing.T) {
	t.Run("invalid", func(t *testing.T) {
		_, err := NewStream("invalid_method", []byte{}, []byte{}, true)
		require.Error(t, err)
		_, err = NewStream("aes-128-cfb", []byte{}, []byte("01234567890abcdef"), true)
		require.Error(t, err)
		_, err = NewStream("aes-128-cfb", []byte("01234567890abcdef"), []byte{}, true)
		require.Error(t, err)

		_, err = NewStream("rc4-md5", []byte{}, []byte("01234567890abcdef"), true)
		require.Error(t, err)
		_, err = NewStream("rc4-md5", []byte("01234567890abcdef"), []byte{}, true)
		require.Error(t, err)
	})
	t.Run("stream", func(t *testing.T) {
		password := "pass_word"
		key := Evp2Key(password, 32)
		iv := Evp2Key("password", 32)
		plainText := []byte("hello world")
		for _, method := range CipherMethods() {
			require.True(t, Valid(method, password))
			require.True(t, HasCipherMethod(method))

			wr, err := NewStream(method, key, iv, true)
			require.NoError(t, err)
			rd, err := NewStream(method, key, iv, false)
			require.NoError(t, err)

			// encrypt
			cipherText := make([]byte, len(plainText))
			wr.XORKeyStream(cipherText, plainText)
			// decrypt
			got := make([]byte, len(cipherText))
			rd.XORKeyStream(got, cipherText)

			assert.Equal(t, plainText, got, fmt.Errorf("method: %s", method))
		}
	})
}
