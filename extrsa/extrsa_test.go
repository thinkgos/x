package extrsa

import (
	"crypto/rand"
	"crypto/rsa"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRsa(t *testing.T) {
	t.Run("", func(t *testing.T) {
		plainText := []byte("hello word github.com/stretchr/testify/require")
		priFilename := "pri.pem"
		pubFilename := "pub.pem"

		pri, pub, err := GenerateKeys(4096)
		require.NoError(t, err)

		err = WriteFile(priFilename, pri)
		require.NoError(t, err)
		defer os.Remove(priFilename)
		err = WriteFile(pubFilename, pub)
		require.NoError(t, err)
		defer os.Remove(pubFilename)

		privateKey, err := LoadPrivateKey(priFilename)
		require.NoError(t, err)
		publicKey, err := LoadPublicKey(pubFilename)
		require.NoError(t, err)

		cipherTest, err := rsa.EncryptPKCS1v15(rand.Reader, publicKey, plainText)
		require.NoError(t, err)
		got, err := rsa.DecryptPKCS1v15(rand.Reader, privateKey, cipherTest)
		require.NoError(t, err)
		assert.Equal(t, plainText, got)
	})
	t.Run("invalid bits", func(t *testing.T) {
		_, _, err := GenerateKeys(1)
		require.Error(t, err)
	})
	t.Run("invalid file", func(t *testing.T) {
		_, err := LoadPrivateKey("invalid")
		require.Error(t, err)
		_, err = LoadPublicKey("invalid")
		require.Error(t, err)
	})
	t.Run("invalid public key", func(t *testing.T) {
		_, err := ParsePublicKey([]byte{})
		require.Error(t, err)
		_, err = ParsePrivateKey([]byte{})
		require.Error(t, err)
	})
}
