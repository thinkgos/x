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
	"crypto/aes"
	"crypto/cipher"
	"crypto/des"
	"crypto/md5"
	"crypto/rc4"
	"encoding/binary"

	"golang.org/x/crypto/blowfish"
	"golang.org/x/crypto/cast5"
	"golang.org/x/crypto/chacha20"
	"golang.org/x/crypto/salsa20/salsa"
	"golang.org/x/crypto/tea"
	"golang.org/x/crypto/twofish"
	"golang.org/x/crypto/xtea"
)

type encDec struct {
	key       []byte
	iv        []byte
	newCipher func(key []byte) (cipher.Block, error)
	newStream func(block cipher.Block, iv []byte) cipher.Stream
}

// complexCipher cipher information
type complexCipher struct {
	keyLen     int
	ivLen      int
	newCipher  func(key []byte) (cipher.Block, error)
	newEncrypt func(block cipher.Block, iv []byte) cipher.Stream
	newDecrypt func(block cipher.Block, iv []byte) cipher.Stream
	newStream  func(*encDec) (cipher.Stream, error)
}

// KeyLen return key len
func (sf complexCipher) KeyLen() int { return sf.keyLen }

// IvLen return iv len
func (sf complexCipher) IvLen() int { return sf.ivLen }

// simpleCiphers cipher information
type simpleCipher struct {
	keyLen    int
	ivLen     int
	newStream func(key, iv []byte) (cipher.Stream, error)
}

// KeyLen return key len
func (sf simpleCipher) KeyLen() int { return sf.keyLen }

// IvLen return iv len
func (sf simpleCipher) IvLen() int { return sf.ivLen }

var complexCiphers = map[string]complexCipher{
	"aes-128-cfb": {16, 16, aes.NewCipher, cipher.NewCFBEncrypter, cipher.NewCFBDecrypter, newStreamWithCipher},
	"aes-192-cfb": {24, 16, aes.NewCipher, cipher.NewCFBEncrypter, cipher.NewCFBDecrypter, newStreamWithCipher},
	"aes-256-cfb": {32, 16, aes.NewCipher, cipher.NewCFBEncrypter, cipher.NewCFBDecrypter, newStreamWithCipher},
	"aes-128-ctr": {16, 16, aes.NewCipher, cipher.NewCTR, cipher.NewCTR, newStreamWithCipher},
	"aes-192-ctr": {24, 16, aes.NewCipher, cipher.NewCTR, cipher.NewCTR, newStreamWithCipher},
	"aes-256-ctr": {32, 16, aes.NewCipher, cipher.NewCTR, cipher.NewCTR, newStreamWithCipher},
	"aes-128-ofb": {16, 16, aes.NewCipher, cipher.NewOFB, cipher.NewOFB, newStreamWithCipher},
	"aes-192-ofb": {24, 16, aes.NewCipher, cipher.NewOFB, cipher.NewOFB, newStreamWithCipher},
	"aes-256-ofb": {32, 16, aes.NewCipher, cipher.NewOFB, cipher.NewOFB, newStreamWithCipher},
	"des-cfb":     {8, 8, des.NewCipher, cipher.NewCFBEncrypter, cipher.NewCFBDecrypter, newStreamWithCipher},
	"des-ctr":     {8, 8, des.NewCipher, cipher.NewCTR, cipher.NewCTR, newStreamWithCipher},
	"des-ofb":     {8, 8, des.NewCipher, cipher.NewOFB, cipher.NewOFB, newStreamWithCipher},
	"3des-cfb":    {24, 8, des.NewTripleDESCipher, cipher.NewCFBEncrypter, cipher.NewCFBDecrypter, newStreamWithCipher},
	"3des-ctr":    {24, 8, des.NewTripleDESCipher, cipher.NewCTR, cipher.NewCTR, newStreamWithCipher},
	"3des-ofb":    {24, 8, des.NewTripleDESCipher, cipher.NewOFB, cipher.NewOFB, newStreamWithCipher},
	"blowfish-cfb": {
		16, 8,
		func(k []byte) (cipher.Block, error) { return blowfish.NewCipher(k) },
		cipher.NewCFBEncrypter, cipher.NewCFBDecrypter, newStreamWithCipher},
	"blowfish-ctr": {
		16, 8,
		func(k []byte) (cipher.Block, error) { return blowfish.NewCipher(k) },
		cipher.NewCTR, cipher.NewCTR, newStreamWithCipher,
	},
	"blowfish-ofb": {
		16, 8,
		func(k []byte) (cipher.Block, error) { return blowfish.NewCipher(k) },
		cipher.NewOFB, cipher.NewOFB, newStreamWithCipher,
	},
	"cast5-cfb": {
		16, 8,
		func(k []byte) (cipher.Block, error) { return cast5.NewCipher(k) },
		cipher.NewCFBEncrypter, cipher.NewCFBDecrypter, newStreamWithCipher,
	},
	"cast5-ctr": {
		16, 8,
		func(k []byte) (cipher.Block, error) { return cast5.NewCipher(k) },
		cipher.NewCTR, cipher.NewCTR, newStreamWithCipher,
	},
	"cast5-ofb": {
		16, 8,
		func(k []byte) (cipher.Block, error) { return cast5.NewCipher(k) },
		cipher.NewOFB, cipher.NewOFB, newStreamWithCipher,
	},
	"twofish-128-cfb": {
		16, 16,
		func(k []byte) (cipher.Block, error) { return twofish.NewCipher(k) },
		cipher.NewCFBEncrypter, cipher.NewCFBDecrypter, newStreamWithCipher,
	},
	"twofish-192-cfb": {
		24, 16,
		func(k []byte) (cipher.Block, error) { return twofish.NewCipher(k) },
		cipher.NewCFBEncrypter, cipher.NewCFBDecrypter, newStreamWithCipher,
	},
	"twofish-256-cfb": {
		32, 16,
		func(k []byte) (cipher.Block, error) { return twofish.NewCipher(k) },
		cipher.NewCFBEncrypter, cipher.NewCFBDecrypter, newStreamWithCipher,
	},
	"twofish-128-ctr": {
		16, 16,
		func(k []byte) (cipher.Block, error) { return twofish.NewCipher(k) },
		cipher.NewCTR, cipher.NewCTR, newStreamWithCipher,
	},
	"twofish-192-ctr": {
		24, 16,
		func(k []byte) (cipher.Block, error) { return twofish.NewCipher(k) },
		cipher.NewCTR, cipher.NewCTR, newStreamWithCipher,
	},
	"twofish-256-ctr": {
		32, 16,
		func(k []byte) (cipher.Block, error) { return twofish.NewCipher(k) },
		cipher.NewCTR, cipher.NewCTR, newStreamWithCipher,
	},
	"twofish-128-ofb": {
		16, 16,
		func(k []byte) (cipher.Block, error) { return twofish.NewCipher(k) },
		cipher.NewOFB, cipher.NewOFB, newStreamWithCipher,
	},
	"twofish-192-ofb": {
		24, 16,
		func(k []byte) (cipher.Block, error) { return twofish.NewCipher(k) },
		cipher.NewOFB, cipher.NewOFB, newStreamWithCipher,
	},
	"twofish-256-ofb": {
		32, 16,
		func(k []byte) (cipher.Block, error) { return twofish.NewCipher(k) },
		cipher.NewOFB, cipher.NewOFB, newStreamWithCipher,
	},
	"xtea-cfb": {16, 8,
		func(k []byte) (cipher.Block, error) { return xtea.NewCipher(k) },
		cipher.NewCFBEncrypter, cipher.NewCFBDecrypter, newStreamWithCipher,
	},
	"xtea-ctr": {16, 8,
		func(k []byte) (cipher.Block, error) { return xtea.NewCipher(k) },
		cipher.NewCTR, cipher.NewCTR, newStreamWithCipher,
	},
	"xtea-ofb": {16, 8,
		func(k []byte) (cipher.Block, error) { return xtea.NewCipher(k) },
		cipher.NewOFB, cipher.NewOFB, newStreamWithCipher,
	},
	"tea-cfb": {16, 8, tea.NewCipher, cipher.NewCFBEncrypter, cipher.NewCFBDecrypter, newStreamWithCipher},
	"tea-ctr": {16, 8, tea.NewCipher, cipher.NewCTR, cipher.NewCTR, newStreamWithCipher},
	"tea-ofb": {16, 8, tea.NewCipher, cipher.NewOFB, cipher.NewOFB, newStreamWithCipher},
}
var simpleCiphers = map[string]simpleCipher{
	"rc4-md5":       {16, 16, newRc4Md5Stream},
	"rc4-md5-6":     {16, 6, newRc4Md5Stream},
	"chacha20":      {32, 12, newChaCha20Stream},
	"chacha20-ietf": {32, 24, newChaCha20IETFStream},
	"salsa20":       {32, 8, newSalsa20Stream},
}

func newStreamWithCipher(ec *encDec) (cipher.Stream, error) {
	block, err := ec.newCipher(ec.key)
	if err != nil {
		return nil, err
	}
	return ec.newStream(block, ec.iv), nil
}

func newRc4Md5Stream(key, iv []byte) (cipher.Stream, error) {
	h := md5.New()
	h.Write(key) // nolint: errcheck
	h.Write(iv)  // nolint: errcheck
	return rc4.NewCipher(h.Sum(nil))
}

func newChaCha20Stream(key, iv []byte) (cipher.Stream, error) {
	return chacha20.NewUnauthenticatedCipher(key, iv)
}

func newChaCha20IETFStream(key, iv []byte) (cipher.Stream, error) {
	return chacha20.NewUnauthenticatedCipher(key, iv)
}

func newSalsa20Stream(key, iv []byte) (cipher.Stream, error) {
	var c salsaStreamCipher
	copy(c.key[:], key[:32])
	copy(c.nonce[:], iv[:8])
	return &c, nil
}

type salsaStreamCipher struct {
	nonce   [8]byte
	key     [32]byte
	counter int
}

func (c *salsaStreamCipher) XORKeyStream(dst, src []byte) {
	var buf []byte
	padLen := c.counter % 64
	dataSize := len(src) + padLen
	if cap(dst) >= dataSize {
		buf = dst[:dataSize]
	} else {
		buf = make([]byte, dataSize)
	}

	var subNonce [16]byte
	copy(subNonce[:], c.nonce[:])
	binary.LittleEndian.PutUint64(subNonce[len(c.nonce):], uint64(c.counter/64))

	// It's difficult to avoid data copy here. src or dst maybe slice from
	// Conn.Read/Write, which can't have padding.
	copy(buf[padLen:], src)
	salsa.XORKeyStream(buf, buf, &subNonce, &c.key)
	copy(dst, buf[padLen:])

	c.counter += len(src)
}
