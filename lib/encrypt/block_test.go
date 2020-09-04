package encrypt

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"errors"
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

func mockErrorNewCipher([]byte) (cipher.Block, error) {
	return nil, errors.New("mock error new cipher")
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

				assert.Equal(t, aes.BlockSize, blk.BlockSize())

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
	t.Run("invalid cipher", func(t *testing.T) {
		key := make([]byte, 16)
		_, err := io.ReadFull(rand.Reader, key)
		require.NoError(t, err)

		bc := BlockStreamCipher{
			cipher.NewCFBEncrypter,
			cipher.NewCFBDecrypter,
		}
		_, err = bc.New(key, mockErrorNewCipher)
		require.Error(t, err)
	})
	t.Run("invalid iv function or length", func(t *testing.T) {
		plainText := []byte("hello world,this is golang language. welcome")
		key := make([]byte, 16)
		_, err := io.ReadFull(rand.Reader, key)
		require.NoError(t, err)

		bmc := BlockStreamCipher{
			cipher.NewCFBEncrypter,
			cipher.NewCFBDecrypter,
		}
		bc, err := bmc.New(key, aes.NewCipher, WithGenerateIv(func(block cipher.Block) ([]byte, error) {
			return nil, errors.New("invalid iv")
		}))
		require.NoError(t, err)

		_, err = bc.Encrypt(plainText)
		require.Error(t, err)
	})
}

func TestBlockModeCipher(t *testing.T) {
	plainText := []byte("helloworld,this is golang language. welcome")
	t.Run("aes", func(t *testing.T) {
		for _, keySize := range aesKeySizes {
			key := make([]byte, keySize)
			_, err := io.ReadFull(rand.Reader, key)
			require.NoError(t, err)

			bc := BlockModeCipher{
				cipher.NewCBCEncrypter,
				cipher.NewCBCDecrypter,
			}
			blk, err := bc.New(key, aes.NewCipher, WithGenerateIv(RandIV))
			require.NoError(t, err)

			assert.Equal(t, aes.BlockSize, blk.BlockSize())

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
	t.Run("invalid cipher", func(t *testing.T) {
		key := make([]byte, 16)
		_, err := io.ReadFull(rand.Reader, key)
		require.NoError(t, err)

		bc := BlockModeCipher{
			cipher.NewCBCEncrypter,
			cipher.NewCBCDecrypter,
		}
		_, err = bc.New(key, mockErrorNewCipher)
		require.Error(t, err)
	})
	t.Run("invalid iv function or length", func(t *testing.T) {
		key := make([]byte, 16)
		_, err := io.ReadFull(rand.Reader, key)
		require.NoError(t, err)

		bmc := BlockModeCipher{
			cipher.NewCBCEncrypter,
			cipher.NewCBCDecrypter,
		}
		bc, err := bmc.New(key, aes.NewCipher, WithGenerateIv(func(block cipher.Block) ([]byte, error) {
			return nil, errors.New("invalid iv")
		}))
		require.NoError(t, err)

		_, err = bc.Encrypt(plainText)
		require.Error(t, err)
	})
	t.Run("invalid input cipher text", func(t *testing.T) {
		key := make([]byte, 16)
		_, err := io.ReadFull(rand.Reader, key)
		require.NoError(t, err)

		bmc := BlockModeCipher{
			cipher.NewCBCEncrypter,
			cipher.NewCBCDecrypter,
		}
		bc, err := bmc.New(key, aes.NewCipher)
		require.NoError(t, err)

		_, err = bc.Decrypt([]byte{})
		require.Error(t, err)
	})
}
