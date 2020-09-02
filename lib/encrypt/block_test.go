package encrypt

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"io"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var aesKeySizes = []int{16, 24, 32}
var encdec = []struct {
	enc, dec func(block cipher.Block, iv []byte) cipher.Stream
}{
	{cipher.NewCFBEncrypter, cipher.NewCFBDecrypter},
	{cipher.NewCTR, cipher.NewCTR},
	{cipher.NewOFB, cipher.NewOFB},
}

func TestBlockStream(t *testing.T) {
	t.Run("encrypt decrypt", func(t *testing.T) {
		plainText := []byte("hello world,this is golang language. welcome")
		for _, encdec := range encdec {
			for _, keySize := range aesKeySizes {
				key := make([]byte, keySize)
				_, err := io.ReadFull(rand.Reader, key)
				require.NoError(t, err)

				bc := BlockStreamCipher{
					encdec.enc,
					encdec.dec,
				}
				blk, err := bc.New(key, aes.NewCipher, WithNewIv(RandIV))
				require.NoError(t, err)

				cipherText, err := blk.Encrypt(plainText)
				require.NoError(t, err)
				want, err := blk.Decrypt(cipherText)
				require.NoError(t, err)
				assert.Equal(t, want, plainText)

				cipherText, err = blk.Encrypt(plainText)
				require.NoError(t, err)
				want, err = blk.Decrypt(cipherText)
				require.NoError(t, err)
				assert.Equal(t, want, plainText)
			}
		}
	})

	t.Run("invalid cipher text", func(t *testing.T) {
		key := make([]byte, 16)
		_, err := io.ReadFull(rand.Reader, key)
		require.NoError(t, err)

		blk, err := aes.NewCipher(key)
		require.NoError(t, err)

		stream := blockStream{
			block:      blk,
			newEncrypt: cipher.NewCFBEncrypter,
			newDecrypt: cipher.NewCFBDecrypter,
		}

		_, err = stream.Encrypt([]byte{})
		require.EqualError(t, err, ErrInputInvalidLength.Error())

		_, err = stream.Decrypt(key[:len(key)-1])
		require.EqualError(t, err, ErrInputNotMoreABlock.Error())
	})
}

func TestBlockModeCipher(t *testing.T) {
	t.Run("aes", func(t *testing.T) {
		plainText := []byte("helloworld,this is golang language. welcome")
		for _, keySize := range aesKeySizes {
			key := make([]byte, keySize)
			_, err := io.ReadFull(rand.Reader, key)
			require.NoError(t, err)

			bc := BlockModeCipher{
				cipher.NewCBCEncrypter,
				cipher.NewCBCDecrypter,
			}
			blk, err := bc.New(key, aes.NewCipher, WithNewIv(RandIV))
			require.NoError(t, err)

			cipherText, err := blk.Encrypt(plainText)
			require.NoError(t, err)
			want, err := blk.Decrypt(cipherText)
			require.NoError(t, err)
			assert.Equal(t, want, plainText)

			cipherText, err = blk.Encrypt(plainText)
			require.NoError(t, err)
			want, err = blk.Decrypt(cipherText)
			require.NoError(t, err)
			assert.Equal(t, want, plainText)
		}
	})
}
