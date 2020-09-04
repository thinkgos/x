// Copyright [2020] [thinkgos] thinkgo@aliyun.com
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package encrypt

import (
	"bytes"
	"crypto/cipher"
	"crypto/rand"
	"errors"
	"io"
)

// error defined
var (
	ErrInputInvalidLength     = errors.New("encoded message length must be more than zero")
	ErrInputNotMoreABlock     = errors.New("decoded message length must be more than a block size")
	ErrInputNotMultipleBlocks = errors.New("decoded message length must be multiple of block size")
	ErrInvalidIvSize          = errors.New("iv length must equal block size")
	ErrUnPaddingOutOfRange    = errors.New("unPadding out of range")
)

// BlockCrypt block crypt interface
type BlockCrypt interface {
	// BlockSize returns the mode's block size.
	BlockSize() int
	// Encrypt plain text. return iv + cipher text
	Encrypt(plainText []byte) ([]byte, error)
	// Encrypt cipher text(iv + cipher text). plain text.
	Decrypt(cipherText []byte) ([]byte, error)
}

// Apply apply
type Apply interface {
	apply(generateIv func(block cipher.Block) ([]byte, error))
}

// Option option
type Option func(apply Apply)

// WithNewIv with custom generate new iv function
// Deprecated: use WithGenerateIv
func WithNewIv(generateIv func(block cipher.Block) ([]byte, error)) Option {
	return WithGenerateIv(generateIv)
}

// WithGenerateIv with custom generate new iv function
func WithGenerateIv(generateIv func(block cipher.Block) ([]byte, error)) Option {
	return func(apply Apply) {
		apply.apply(generateIv)
	}
}

// RandIV generate rand iv by rand.Reader
func RandIV(block cipher.Block) ([]byte, error) {
	iv := make([]byte, block.BlockSize())
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return nil, err
	}
	return iv, nil
}

// BlockStreamCipher block stream cipher
// support:
// 		cfb: cipher.NewCFBEncrypter, cipher.NewCFBDecrypter
// 		ctr: cipher.NewCTR, cipher.NewCTR
// 		ofb: cipher.NewOFB, cipher.NewOFB
type BlockStreamCipher struct {
	NewEncrypt func(block cipher.Block, iv []byte) cipher.Stream
	NewDecrypt func(block cipher.Block, iv []byte) cipher.Stream
}

// New new with newCipher and key
// newCipher support follow or implement func(key []byte) (cipher.Block, error):
// 		aes
// 		cipher
// 		des
// 		blowfish
// 		cast5
// 		twofish
// 		xtea
// 		tea
func (sf *BlockStreamCipher) New(key []byte,
	newCipher func(key []byte) (cipher.Block, error), opts ...Option) (BlockCrypt, error) {
	block, err := newCipher(key)
	if err != nil {
		return nil, err
	}
	bs := &blockStream{
		block:      block,
		newEncrypt: sf.NewEncrypt,
		newDecrypt: sf.NewDecrypt,
	}
	for _, opt := range opts {
		opt(bs)
	}
	return bs, nil
}

// BlockModeCipher block mode cipher
// support:
//      cbc: cipher.NewCBCEncrypter, cipher.NewCBCDecrypter
type BlockModeCipher struct {
	NewEncrypt func(block cipher.Block, iv []byte) cipher.BlockMode
	NewDecrypt func(block cipher.Block, iv []byte) cipher.BlockMode
}

// New new with newCipher and key
// newCipher support follow or implement func(key []byte) (cipher.Block, error):
// 		aes
// 		cipher
// 		des
// 		blowfish
// 		cast5
// 		twofish
// 		xtea
// 		tea
func (sf *BlockModeCipher) New(key []byte,
	newCipher func(key []byte) (cipher.Block, error), opts ...Option) (BlockCrypt, error) {
	block, err := newCipher(key)
	if err != nil {
		return nil, err
	}
	bb := &blockBlock{
		block:      block,
		newEncrypt: sf.NewEncrypt,
		newDecrypt: sf.NewDecrypt,
	}
	for _, opt := range opts {
		opt(bb)
	}
	return bb, nil
}

type blockStream struct {
	block      cipher.Block
	generateIv func(block cipher.Block) ([]byte, error)
	newEncrypt func(block cipher.Block, iv []byte) cipher.Stream
	newDecrypt func(block cipher.Block, iv []byte) cipher.Stream
}

func (sf *blockStream) apply(generateIv func(block cipher.Block) ([]byte, error)) {
	sf.generateIv = generateIv
}

func (sf *blockStream) BlockSize() int {
	return sf.block.BlockSize()
}

func (sf *blockStream) Encrypt(plainText []byte) ([]byte, error) {
	if len(plainText) == 0 {
		return nil, ErrInputInvalidLength
	}
	blockSize := sf.block.BlockSize()

	ivFunc := RandIV
	if sf.generateIv != nil {
		ivFunc = sf.generateIv
	}

	iv, err := ivFunc(sf.block)
	if err != nil || len(iv) != blockSize {
		return nil, ErrInvalidIvSize
	}

	cipherText := make([]byte, blockSize+len(plainText))
	copy(cipherText[:blockSize], iv)
	sf.newEncrypt(sf.block, iv).XORKeyStream(cipherText[blockSize:], plainText)
	return cipherText, nil
}

func (sf *blockStream) Decrypt(cipherText []byte) ([]byte, error) {
	blockSize := sf.block.BlockSize()
	if len(cipherText) < blockSize {
		return nil, ErrInputNotMoreABlock
	}
	iv, msg := cipherText[:blockSize], cipherText[blockSize:]
	sf.newDecrypt(sf.block, iv).XORKeyStream(msg, msg)
	return msg, nil
}

type blockBlock struct {
	block      cipher.Block
	generateIv func(block cipher.Block) ([]byte, error)
	newEncrypt func(block cipher.Block, iv []byte) cipher.BlockMode
	newDecrypt func(block cipher.Block, iv []byte) cipher.BlockMode
}

func (sf *blockBlock) apply(generateIv func(block cipher.Block) ([]byte, error)) {
	sf.generateIv = generateIv
}

func (sf *blockBlock) BlockSize() int {
	return sf.block.BlockSize()
}

// Encrypt encrypt
func (sf *blockBlock) Encrypt(plainText []byte) ([]byte, error) {
	blockSize := sf.block.BlockSize()

	ivFunc := RandIV
	if sf.generateIv != nil {
		ivFunc = sf.generateIv
	}
	iv, err := ivFunc(sf.block)
	if err != nil || len(iv) != blockSize {
		return nil, ErrInvalidIvSize
	}

	orig := PCKSPadding(plainText, blockSize)
	cipherText := make([]byte, blockSize+len(orig))
	copy(cipherText[:blockSize], iv)
	sf.newEncrypt(sf.block, iv).CryptBlocks(cipherText[blockSize:], orig)
	return cipherText, nil
}

// Decrypt decrypt
func (sf *blockBlock) Decrypt(cipherText []byte) ([]byte, error) {
	blockSize := sf.block.BlockSize()
	if len(cipherText) == 0 || len(cipherText)%blockSize != 0 {
		return nil, ErrInputNotMultipleBlocks
	}
	iv, msg := cipherText[:blockSize], cipherText[blockSize:]
	sf.newDecrypt(sf.block, iv).CryptBlocks(msg, msg)
	return PCKSUnPadding(msg)
}

// PCKSPadding PKCS#5和PKCS#7 填充
func PCKSPadding(origData []byte, blockSize int) []byte {
	padSize := blockSize - len(origData)%blockSize
	padText := bytes.Repeat([]byte{byte(padSize)}, padSize)
	return append(origData, padText...)
}

// PCKSUnPadding PKCS#5和PKCS#7 解填充
func PCKSUnPadding(origData []byte) ([]byte, error) {
	length := len(origData)
	if length == 0 {
		return nil, ErrUnPaddingOutOfRange
	}
	unPadSize := int(origData[length-1])
	if unPadSize > length {
		return nil, ErrUnPaddingOutOfRange
	}
	return origData[:(length - unPadSize)], nil
}
